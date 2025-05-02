.PHONY: drone

drone:
	drone lint
	drone --server https://drone.grafana.net sign --save muandane/differ