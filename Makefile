DAY=$(shell date +%d)

today:
	cp -R 2020/boilerplate 2020/day$(DAY)
	sed -i '' "s/XXX_DATE_XXX/day$(DAY)/" 2020/day$(DAY)/main.go

