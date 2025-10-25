TESTNUM = 01

run:
	go build -o $(PROG).x -compiler gc $(PROG)/$(PROG).go && chmod +x ./$(PROG).x && ./$(PROG).x < $(PROG)/in$(TESTNUM).txt && rm $(PROG).x

generate:
	mkdir -p $(PROG) && cp template/template.go $(PROG)/$(PROG).go && touch $(PROG)/in01.txt
