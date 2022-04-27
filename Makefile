start:
	kind create cluster --config deploy/config.yaml || true
	
	kubectl create -f deploy/k8s/0-crds.yaml || true
	kubectl wait --for condition=established --timeout=60s crd/servicemonitors.monitoring.coreos.com
	kubectl wait --for condition=established --timeout=60s crd/prometheuses.monitoring.coreos.com

	kubectl apply -f deploy/k8s || true
	kubectl create clusterrolebinding default-admin --clusterrole admin --serviceaccount default:default
	kubectl create clusterrolebinding default-metrics --clusterrole system:metrics-server --serviceaccount default:default

	# timestamp 2022-04-22 11:34:08

block:
	go run main.go