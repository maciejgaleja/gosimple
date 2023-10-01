# gosimple
Proof of Concepts and experimentation should be fun. Don’t waste time on implementation, focus on the cool things.

> :warning: This project is not intended for the production use.

The goal is simplicity, even at the expense of optimization and performance. Both interfaces and implementation prioritize simplicity.

## Background
It's safe to say that there are some fundamental building blocks of applications, such as storage, databases, and messaging.

This project offers interfaces and implementations for these components. You can design an app using the provided interfaces and choose the implementation that best suits your needs.

## Example
Let’s say you want to learn new technology, or want to try a new idea, for example backend for a web app. You will probably want to start with local deployment. Go simple. Write your code around these interfaces and select local backend for them - store assets in the filesystem, and data in JSON files. Sure, it’s not optimal, but it’s simple.

If you want to migrate to the cloud or another environment, simply change the implementation. This will likely involve updating just one line of code.

## Features

> :warning: This project is currently under active development. While I do not intend to break compatibility at any point, I cannot guarantee that it will never happen in the future.

| Service | Supported backends | Ideas/planned |
| --- | --- | --- |
| Storage | Local directory <br /> AWS S3 | |
| Key-value store | Local JSON file <br /> NoSQL with primary key | Redis |
| NoSQL with primary key | AWS DynamoDB | Local JSON file |