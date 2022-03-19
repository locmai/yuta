current status: alpha

```
             +----------------+         +------------------+
             |                |         |                  |
             |  chat clients  |         |    nlu clients   |
             |  matrix,slack  |         | diaglogflow,luis |
             |                |         |                  |
             +-------^--------+         +---------^--------+
                     |                            |
                     |                            |
                     |                            |
+---------------+    |                            |
|               |    |                            |
|    common     |    |                            |
|               |    +-----+--------------+-------+
+-------+-------+          |              |
        +------------------>   messaging  +-------+-------+
        |                  |              |       +-------v-----+
        |                  +--------------+       |    nats     |
        |                                         |  jetstream  |
        |                  +--------------+       |             |
        |                  |              |       +------^----^-+
        +------------------>     core     |              |    |
                           |              +--------------+    |
                           +-------+------+                   |
                                   |                          |
                                   |                          |
                                   |     +----------------+   |
                                   |     |                |   |
                                   +----->    kubeops     |   |
                                   |     |                |   |
                                   |     +----------------+   |
                                   |                          |
                                   |     +----------------+   |
                                   |     |                |   |
                                   +----->     argocd     |   |
                                   |     |                |   |
                                   |     +----------------+   |
                                   |                          |
                                   +-----> ...                |
                                                              |
+--------------+           +--------------+                   |
|              |           |              |                   |
|  prometheus  +----------->     pad      |                   |      
|              |           |              +-------------------+
+-------+------+           +-------+------+
```

- common: /common setup, code base, infra
  - [x] config
  - [x] metrics
  - [x] nats
  - [ ] tracing
  - [x] utils
- messaging:
  - chat clients:
    - [x] matrix
    - [ ] slack
  - natural language understanding clients:
    - [x] diaglogflow
    - [ ] luis
- core:
  - appservices:
    - [x] kubeops
    - [ ] github
    - [ ] argocd
    - [ ] faq
- pad: tbd
- others:
  - [x] dockerfile
  - [x] chart
  - [ ] lint
  - [ ] test
  - [ ] docs

