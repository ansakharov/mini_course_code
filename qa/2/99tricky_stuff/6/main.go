package main

import (
	"context"
	"io"
)

//Пример для разбора // 5. // Есть функция processDataInternal, которая может выполняться неопределенно долго.
//Чтобы контролировать процесс, мы добавили таймаут выполнения ф-ии через context.
//Какие недостатки кода ниже?
func (s *Service) ProcessData(timeoutCtx context.Context, r io.Reader) error {
	errCh := make(chan error)
	go func() { errCh <- s.processDataInternal(r) }()
	select {
	case err := <-errCh:
		return err
	case <-timeoutCtx.Done():
		return timeoutCtx.Err()
	}
}