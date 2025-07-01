push:
	git add .
	@if [ -z "$(m)" ]; then msg="update"; else msg="$(m)"; fi; git commit -m "$$msg"
	git push
	@echo "Pushed changes to remote repository."