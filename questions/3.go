package main

import "sync"

//В языке Go существуют два основных типа блокировок для управления доступом к общим данным: `sync.Mutex`
//и `sync.RWMutex`.
//
//1. **sync.Mutex:**
//    - `sync.Mutex` используется для создания критической секции, в которой только один поток может получить доступ к
//      общим данным в определенный момент времени.
//    - Это простая блокировка, которая обеспечивает эксклюзивный доступ к общим данным.
//    - При вызове `Lock()` блокировка ожидает, пока не будет доступна, и затем захватывает ее. Другие потоки будут
//      ожидать, пока блокировка не будет освобождена с помощью `Unlock()`.
//    - `Mutex` может быть захвачена только одним потоком. Если другой поток пытается захватить блокировку, пока она уже
//      захвачена, он будет блокирован до тех пор, пока блокировка не будет освобождена.
//
//2. **sync.RWMutex:**
//    - `sync.RWMutex` (Read-Write Mutex) также используется для управления доступом к общим данным, но предоставляет
//      более гибкую семантику блокировки.
//    - `RWMutex` позволяет нескольким потокам читать общие данные одновременно, но позволяет только одному потоку
//      записывать (или изменять) данные в определенный момент времени.
//    - При вызове `RLock()` блокировка типа `RWMutex` разрешает множественное чтение данных, пока не будет
//      вызван `RUnlock()`.
//    - Когда блокировка типа `RWMutex` захвачена для чтения, другие потоки могут продолжать читать, но не могут
//      записывать. Когда блокировка захвачена для записи (`Lock()`), другие потоки не могут получить доступ к общим
//      данным до тех пор, пока блокировка не будет освобождена.

func main() {
	mu := sync.Mutex{}
	mu.Lock()
	mu.Unlock()
	mu.TryLock()

	rwmu := sync.RWMutex{}
	rwmu.Lock()
	rwmu.Unlock()
	rwmu.TryLock()

	rwmu.RLock()
	rwmu.RUnlock()
	rwmu.TryRLock()

	rwmu.RLocker()
}