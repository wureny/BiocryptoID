# Technical Highlights

## Back-end

### High Performance

Go's concurrency model and garbage collection mechanism enable it to perform well in handling high-concurrency requests.&#x20;

### Powerful Retry Strategy&#x20;

We have implemented a flexible and powerful retry mechanism with the following main features:&#x20;

#### Strategy interface design

Through the Strategy interface, different retry strategies can be easily implemented and switched.

#### Exponential Backoff Algorithm

ExponentialStrategy implements the classic exponential backoff algorithm to effectively prevent system overload.&#x20;

#### Adaptive Adjustment

The retry interval increases with the number of attempts, minimizing the pressure on the system.

#### Jitter Mechanism

Random jitter is introduced through the MaxJitter parameter to avoid the "herd effect" caused by multiple clients retrying at the same time.&#x20;

#### Configurability

Supports setting the minimum and maximum retry intervals to adapt to different application scenarios.&#x20;

#### Fixed interval strategy

FixedStrategy provides a simple fixed interval retry option, suitable for specific scenarios



## Front-end

### React framework

provides efficient user interface development and good user experience.

### Responsive design

ensures smoothness.

## AWS Service

* AWS Lambda and API gateway enables serverless backend processing&#x20;
* S3 secure storage ensures data privacy and integrity



## Biometric processing

AWS Rekognition is used for biometric processing.

### Advanced deep learning algorithm

Able to accurately identify facial features under various lighting conditions and angles Diverse functions.

### Face matching

Able to compare the similarity of faces in two pictures, suitable for identity authentication

### Attribute detection

Analyze facial attributes such as age range, gender, and emotions.





## Ethereum smart contract management DID&#x20;

### Identity management

DID is a unique identifier with decentralized immutability (guaranteed by blockchain characteristics), which is important in long-term identity management&#x20;

### Cross-platform

interoperability, no need for re-identification

### scalability

Assuming that we have data a and can provide function b, through did, any platform can reference our function b, while not knowing our data a &#x20;

### privacy protection

With did, within a certain period of time (such as a week). We can use it instead of facial information to make verification requests, without frequent face scanning, and improve privacy protection capabilities



## Integrated zkp service&#x20;

Zero-knowledge proof can prove the truth of something to a third party without exposing any knowledge Data is minimized, and the verifier only obtains necessary information and does not access redundant personal data. ZK proofs are usually one-time, which prevents verification data from being collected and reused.

