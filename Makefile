production = quiet-meadow-2369


.PHONY: deploy-production
deploy-production:
	fly deploy --app $(production)
