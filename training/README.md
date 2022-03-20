# Training data for RASA NLU server

Install the required packages (Python 3.8.13 required):

```
virtualenv venv
source ./venv/bin/activate
pip install -r requirements.txt
```

To train the NLU model:

```
rasa train nlu
```

To run the NLU server with the trained model:

```
rasa run --enable-api -m models/<model_name>.tar.gz
```

Simple curl test:

```
curl localhost:5005/model/parse -d '{"text":"scale the helloworld deployment in the default namespace to 15 please"}'
```

Sample result:

```json
{"text":"scale the helloworld deployment in the default namespace to 15 please","intent":{"name":"scale_deployment","confidence":1.0},"entities":[{"entity":"name","start":10,"end":20,"confidence_entity":0.9994369149208069,"value":"helloworld","extractor":"DIETClassifier"},{"entity":"namespace","start":39,"end":46,"confidence_entity":0.9992471933364868,"value":"default","extractor":"DIETClassifier"},{"entity":"replicas","start":60,"end":62,"confidence_entity":0.9995238780975342,"value":"15","extractor":"DIETClassifier"}],"text_tokens":[[0,5],[6,9],[10,20],[21,31],[32,34],[35,38],[39,46],[47,56],[57,59],[60,62],[63,69]],"intent_ranking":[{"name":"scale_deployment","confidence":1.0},{"name":"goodbye","confidence":2.7406880320768323e-8},{"name":"greet","confidence":8.079714675091054e-9}],"response_selector":{"all_retrieval_intents":[],"default":{"response":{"responses":null,"confidence":0.0,"intent_response_key":null,"utter_action":"utter_None"},"ranking":[]}}
```
