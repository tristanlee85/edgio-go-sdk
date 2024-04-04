# Changelog

## [1.0.0-beta.5](https://github.com/42dx/edgio-go-sdk/compare/v1.0.0-beta.4...v1.0.0-beta.5) (2024-04-04)

### ⚙️ New Features

* **env-var**: added GetByKey func to edgio/variable pkg [[#42](https://github.com/42dx/edgio-go-sdk/issues/42)] ([29296cc](https://github.com/42dx/edgio-go-sdk/commit/29296ccba7c6048b86ce5d2e21f085523b549287))
* **env**: added GetByName func to edgio/env pkg [[#42](https://github.com/42dx/edgio-go-sdk/issues/42)] ([93c2894](https://github.com/42dx/edgio-go-sdk/commit/93c28944f7796ca538dd1c7c0fb8e999c0b2da16))
* **property**: added GetBySlug func to edgio/property pkg [[#42](https://github.com/42dx/edgio-go-sdk/issues/42)] ([f5c87c5](https://github.com/42dx/edgio-go-sdk/commit/f5c87c5bbf853b952f04fd4a24057fff57f54a58))
* **util**: added base GetByAttr util func [[#42](https://github.com/42dx/edgio-go-sdk/issues/42)] ([62605bb](https://github.com/42dx/edgio-go-sdk/commit/62605bb430fa5e6a53ba0022aac1bbb26e602ecd))

## [1.0.0-beta.4](https://github.com/42dx/edgio-go-sdk/compare/v1.0.0-beta.3...v1.0.0-beta.4) (2024-04-02)

### ⚙️ New Features

* **env-var**: implemented FilterList on variable pkg [[#25](https://github.com/42dx/edgio-go-sdk/issues/25)] ([5a96ae4](https://github.com/42dx/edgio-go-sdk/commit/5a96ae47df30d26060f315a37f47ceb097322228))
* **env**: implemented FilterList on env pkg [[#25](https://github.com/42dx/edgio-go-sdk/issues/25)] ([34fad62](https://github.com/42dx/edgio-go-sdk/commit/34fad62dc51773178bc72175d18ee0d8f8411f74))
* **property**: implemented FilterList on property pkg [[#25](https://github.com/42dx/edgio-go-sdk/issues/25)] ([a755d4a](https://github.com/42dx/edgio-go-sdk/commit/a755d4a625a94054a04ab397cae87b03ef93815c))
* **util**: added generic filter list func [[#25](https://github.com/42dx/edgio-go-sdk/issues/25)] ([1d7b413](https://github.com/42dx/edgio-go-sdk/commit/1d7b413bcba173aef1509a768a29f55c17126c92))

## [1.0.0-beta.3](https://github.com/42dx/edgio-go-sdk/compare/v1.0.0-beta.2...v1.0.0-beta.3) (2024-03-15)

### ⚙️ New Features

* **env-var**: added env-var client and List method [[#20](https://github.com/42dx/edgio-go-sdk/issues/20)] ([d75f9c6](https://github.com/42dx/edgio-go-sdk/commit/d75f9c6e74a7be61457e1a6db76e54b15768927b))

## [1.0.0-beta.2](https://github.com/42dx/edgio-go-sdk/compare/v1.0.0-beta.1...v1.0.0-beta.2) (2024-03-14)

### ⚙️ New Features

- **env**: added list method to env client [[#15](https://github.com/42dx/edgio-go-sdk/issues/15)] ([4d0b5b6](https://github.com/42dx/edgio-go-sdk/commit/4d0b5b6b74897d599a2782f46a553105b63607a7))
- **env**: added env client and base structs [[#15](https://github.com/42dx/edgio-go-sdk/issues/15)] ([5c8d186](https://github.com/42dx/edgio-go-sdk/commit/5c8d1869322682b5491322895801a145cd2b1bea))
- **property**: added list method to property client [[#10](https://github.com/42dx/edgio-go-sdk/issues/10)] ([0dfae98](https://github.com/42dx/edgio-go-sdk/commit/0dfae98410a6c3a8b009247eb883e7c564d22ac4))
- **property**: added property client and base structs [[#10](https://github.com/42dx/edgio-go-sdk/issues/10)] ([0dfae98](https://github.com/42dx/edgio-go-sdk/commit/0dfae98410a6c3a8b009247eb883e7c564d22ac4))

## 1.0.0-beta.1 (2024-03-01)

### ⚙️ New Features

- **client**: added singleton edgio client ([86fab54](https://github.com/42dx/edgio-go-sdk/commit/86fab5410ec2bbbc2e0bf77ad92cdc46f731c63b))
- **client**: edgio central package piece ([1a831bf](https://github.com/42dx/edgio-go-sdk/commit/1a831bfbf135f831f9f17cc17ea1f70b15c86a4d))
- **common**: adding common package to avoid cyclic imports ([f965994](https://github.com/42dx/edgio-go-sdk/commit/f965994ef5f627b34a0a95f3502d6c6a916fd227))
- **init**: create go module ([0ff1e7e](https://github.com/42dx/edgio-go-sdk/commit/0ff1e7ebf86297b8d11959cb2b0fb0d855abd6e0))
- **org**: added org package (get method) ([249634e](https://github.com/42dx/edgio-go-sdk/commit/249634e07c76fbbfe3d5abe182ddfb17fc685874))
- **token**: get access token func first scaffold ([4519759](https://github.com/42dx/edgio-go-sdk/commit/4519759267fe7abaaf09d19daff0bac3ed095513))
- **token**: moving token calls to internal ([fe9894c](https://github.com/42dx/edgio-go-sdk/commit/fe9894c45ee731570686081458c45a29e80cc3e9))
- **utils**: adding http internal utils ([7e99044](https://github.com/42dx/edgio-go-sdk/commit/7e99044cac89b3d8aca6f4a06393195560456244))
