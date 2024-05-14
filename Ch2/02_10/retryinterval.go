package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// vamos a reintentar una captura en tres situaciones diferentes:
// - la primera asegurando que siempre falle
// - la segunda fallando pero porque la función de captura agota el tiempo de espera.
// - y la última, que siempre se consiga.
func retryCapture() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("=========== 1: siempre falla")
	retryCaptureWithInterval(ctx, time.Second, func(ctx context.Context) error {
		// esta función siempre fallará la captura
		return errors.New("la captura ha fallado")
	})

	fmt.Println("=========== 2: falla porque la captura es muy lenta")
	retryCaptureWithInterval(ctx, time.Second, func(ctx context.Context) error {
		// 3 segundos para agotar el tiempo de espera en la segunda iteración
		time.Sleep(3 * time.Second)
		return fmt.Errorf("tiempo de espera agotado")
	})

	fmt.Println("=========== 3: siempre se consigue")
	retryCaptureWithInterval(ctx, time.Second, func(ctx context.Context) error {
		// esta función siempre conseguirá la captura
		return nil
	})
}

// retryCaptureWithInterval: el contexto definirá el posible tiempo de expiración, e intentaremos la captura
// con cierta frequencia, definida por retryInterval. La función captura es la que realizará la misma,
// y podrá lanzar un error para indicar que ha fallado.
func retryCaptureWithInterval(ctx context.Context, retryInterval time.Duration, captureFn func(ctx context.Context) error) {
	for {
		// vamos a probar a capturar un pokemon
		err := captureFn(ctx)
		if err == nil {
			fmt.Println("pokemon capturado con éxito")
			return
		}

		// comprobación de que el tiempo de espera haya expirado
		if ctx.Err() != nil {
			fmt.Println("el tiempo de espera ha expirado (1):", ctx.Err())
			return
		}

		// vamos a intentarlo de nuevo pasado un tiempo de reintento
		fmt.Printf("wait %s before trying again: %s\n", retryInterval, err)

		// timer nos va a ayudar con los reintentos. Utiliza un channel en su estructura interna
		// para gestionarlos.
		t := time.NewTimer(retryInterval)

		select {
		case <-ctx.Done():
			fmt.Println("el tiempo de espera ha expirado (2):", ctx.Err())
			t.Stop()
			return
		case <-t.C:
			fmt.Println("intentando la captura de nuevo")
		}
	}
}
