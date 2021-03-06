run:
	kubectl apply -f ./infra/k8s
stop:
	kubectl delete service auth-srv \
	&& kubectl delete deployment auth-depl \
	&& kubectl delete ingress ingress-service \
	&& kubectl delete service auth-mongo-srv \
	&& kubectl delete deployment auth-mongo-depl \
	&& kubectl delete service client-srv \
	&& kubectl delete deployment client-depl \
	&& kubectl delete service tickets-srv \
	&& kubectl delete deployment tickets-depl \
	&& kubectl delete service tickets-mongo-srv \
	&& kubectl delete deployment tickets-mongo-depl \
	&& kubectl delete service nats-srv \
	&& kubectl delete deployment nats-depl \
	&& kubectl delete service orders-mongo-srv \
	&& kubectl delete deployment orders-mongo-depl \
	&& kubectl delete service orders-srv \
	&& kubectl delete deployment orders-depl