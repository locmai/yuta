train:
	rasa train -d ./domain.yml --fixed-model-name yutaops-dev

dev:
	rasa run --enable-api -m models/yutaops-dev.tar.gz

test:
	rasa test -m models/yutaops-dev.tar.gz
