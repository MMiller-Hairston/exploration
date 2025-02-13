# CI

### What is this?

This is an exploration in effective distributed systems using messaging. The overall project will be split into two major offerings, the CI system (priority #1) and the client app (priority #2).

### Offerings & Components

**CI System**

- **CI Server**: Will be the main entry point of the CI system. Will take events from the SCM system and publish messages to the queue for triggering workflows.

- **Message Queue**: Will be responsible for getting events published by the server and queing them up for consumption by the workers. Leaning towards Kafka and it's at most once delivery. Once a worker takes on the message it will be it's responsibility to report the status the job.

- **Workers**: Will be responsible for fetching jobs from the queue when ready and running those jobs to completion.

- **Log Storage**: Will store job results for viewing after a worker has moved on to the next job. Will likely be MongoDb or S3.

**Client App**

- **UI**: User interface for configuring workflows, governance, jobs, and viewing results.

- **App Server**: Will surface stored data to the UI.

- **App Storage** - Will store data for the client app. Possibly MongoDb since it can store both the app data and the job logs.
