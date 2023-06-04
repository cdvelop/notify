package notify

import "github.com/cdvelop/model"

type publisher interface {
	MainName() string // nombre objeto publicador
}

type listen_to interface {
	MainName() string           // nombre objeto que escucha
	SubscriptionList() []string // listado de suscripciones
	subscription                // función receptora de solicitudes
}

type subscription interface {
	NotificationReceiver(rq *model.Request, out chan<- *model.Request)
}
