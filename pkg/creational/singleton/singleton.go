package singleton

import (
	"sync"
	"sync/atomic"
)

// nolint:gochecknoglobals
var (
	mu          sync.Mutex
	initialized uint32
	once        sync.Once

	instances = make(map[string]*chocolateBoiler)
)

// GetFullInstance Получить экземпляр.
// Простая синхронизация через mutex.
// Не самое удачно решение, если необходима высокая производительность.
// Данный метод станет "бутылочным горлышком", так как в один момент времени только одна go-рутина
// сможет получить доступ к экземпляру, остальные будут ждать.
func GetFullInstance() ChocolateBoiler {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := instances["full"]; !ok {
		instances["full"] = &chocolateBoiler{empty: true}
	}

	return instances["full"]
}

// GetConditionalInstance Получить экземпляр.
// Более удачное решение, так как не происходит блокировка при каждом обращении, а только при первом.
// Дополнительная проверка нужна на случай, если произойдет несколько обращений одновременно, и, как следствие,
// несколько go-рутин пройдут первую проверку. Дальше go-рутины будут по очереди устанавливать блокировку и снова
// проверять, был ли уже проинициализирован экземпляр.
func GetConditionalInstance() ChocolateBoiler {
	if _, ok := instances["conditional"]; !ok {
		mu.Lock()
		defer mu.Unlock()

		if instances["conditional"] == nil {
			instances["conditional"] = &chocolateBoiler{empty: true}
		}
	}

	return instances["conditional"]
}

// GetAtomicInstance Получить экземпляр.
// Аналогично GetConditionalInstance, но используется более надежный метод проверки был ли экземпляр уже
// проинициализирован. Из-за оптимизации при компиляции нет никакой уверенности, что проверка экземпляра
// будет выполняться атомарно, поэтому используется специальный пакет atomic, гарантирующий атомарность при
// установке и проверки ключа.
func GetAtomicInstance() ChocolateBoiler {
	if atomic.LoadUint32(&initialized) == 0 {
		mu.Lock()
		defer mu.Unlock()

		if initialized == 0 {
			instances["atomic"] = &chocolateBoiler{empty: true}

			atomic.StoreUint32(&initialized, 1)
		}
	}

	return instances["atomic"]
}

// GetOnceInstance Получить экземпляр.
// Самый оптимальный вариант. Здесь используется объект Once из стандартного пакета sync.
// Once это объект, который позволяет выполнять некоторое действие только один раз.
//
// type Once struct {
//    m    Mutex
//    done uint32
// }
//
// Do вызывает функцию f только в том случае, если это первый вызов Do для
// этого экземпляра Once. Другими словами, если у нас есть var once Once и
// once.Do(f) будет вызываться несколько раз, f выполнится только в
// момент первого вызова, даже если f будет иметь каждый раз другое значение.
// Для вызова нескольких функций таким способом нужно несколько
// экземпляров Once.
//
// Do предназначен для инициализации, которая должна выполняться единожды.
// Так как f ничего не возвращает, может быть необходимым использование
// замыкания для передачи параметров в функцию, выполняемую Do:
// config.once.Do(func() { config.init(filename) })
//
// Поскольку ни один вызов к Do не завершится пока не произойдет
// первый вызов f, то f может заблокировать последующие вызовы
// Do и получится deadlock.
//
// Если f паникует, то Do считает это обычным вызовом и, при последующих
// вызовах, Do не будет вызывать f.
//
// func (o *Once) Do(f func()) {
//    if atomic.LoadUint32(&o.done) == 1 { // Check
//        return
//    }
//
//    o.m.Lock()                           // Lock
//    defer o.m.Unlock()
//
//    if o.done == 0 {                     // Check
//        defer atomic.StoreUint32(&o.done, 1)
//        f()
//    }
// }.
func GetOnceInstance() ChocolateBoiler {
	once.Do(func() {
		instances["once"] = &chocolateBoiler{empty: true}
	})

	return instances["once"]
}
