package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/overview", getRoot)
	mux.HandleFunc("/hello", getHello)

	// Mux configuration, 2 servers
	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr:    ":8000",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().string())
			return ctx
		},

		serverTwo := &http.Server {
			Addr: ":4444",
			Handler: mux,
			baseContext: func(l net.Listener) context.Context {
				ctx = context.WithValue(ctx, keyServerAddr, l.Addr().string())
				return ctx
			},
		}
	}

	// Error Handling - Add in more error functions if needed
	err := http.ListenAndServe(":8000", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if errors.Is(err, http.ErrHandlerTimeout) {
		fmt.Printf("handler timeout occurred")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got / OR /overview request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "This is my app!\n") // print to client
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Hello, HTTP!\n")
}
