# Changelog

## [1.5.1](https://github.com/Aimerny/kook-go/compare/v1.5.0...v1.5.1) (2024-09-04)


### Bug Fixes

* :bug: Fix potential concurrency issue in EventQueue's ([#35](https://github.com/Aimerny/kook-go/issues/35)) ([924b4f7](https://github.com/Aimerny/kook-go/commit/924b4f776b11e4ec82d7199e36f7de998731e3c0))
* 🐛 Heartbeat scheduled task repeats during reconnection ([#38](https://github.com/Aimerny/kook-go/issues/38)) ([840398c](https://github.com/Aimerny/kook-go/commit/840398ce6ef820481023966bf9859b9f8360b9ca))

## [1.5.0](https://github.com/Aimerny/kook-go/compare/v1.4.0...v1.5.0) (2024-08-24)


### Features

* :sparkles: improve some features of the card module ([1adcc61](https://github.com/Aimerny/kook-go/commit/1adcc61903dfcc5a4cad4a815f8a2b237d0b836e))

## [1.4.0](https://github.com/Aimerny/kook-go/compare/v1.3.3...v1.4.0) (2024-08-06)


### Features

* :sparkles: add asset create action && add post custom header ([#29](https://github.com/Aimerny/kook-go/issues/29)) ([6a65e7d](https://github.com/Aimerny/kook-go/commit/6a65e7d85bc3f2de7cb2643caa2281a5f70c1a96))

## [1.3.3](https://github.com/Aimerny/kook-go/compare/v1.3.2...v1.3.3) (2024-07-22)


### Bug Fixes

* :bug: fix ws connect([#26](https://github.com/Aimerny/kook-go/issues/26)) ([09b7dce](https://github.com/Aimerny/kook-go/commit/09b7dce7bf2688c0b6dcbb5b3d07137120d9b550))

## [1.3.2](https://github.com/Aimerny/kook-go/compare/v1.3.1...v1.3.2) (2024-07-17)


### Bug Fixes

* :bug: fix websocket connection problem ([476d776](https://github.com/Aimerny/kook-go/commit/476d7768cd46eb80bf3bc666bc5e6484a884566b)), closes [#25](https://github.com/Aimerny/kook-go/issues/25)
* :bug: get gateway continous ([82a39aa](https://github.com/Aimerny/kook-go/commit/82a39aa474660d42ac49b2b283df5ad0b356e3b0))

## [1.3.1](https://github.com/Aimerny/kook-go/compare/v1.3.0...v1.3.1) (2024-07-11)


### Bug Fixes

* :bug: fix bot panic when network has broken ([2113c13](https://github.com/Aimerny/kook-go/commit/2113c13b7fa4191fdccf4a5fb639fefcc2618c18))

## [1.3.0](https://github.com/Aimerny/kook-go/compare/v1.2.2...v1.3.0) (2024-07-10)


### Features

* :tada:增加Post请求 ([12e6d66](https://github.com/Aimerny/kook-go/commit/12e6d669e8a7db7a90d42ce169d96b7570418912))
* ✨ 更新消息发送与更新消息功能 ([20608a6](https://github.com/Aimerny/kook-go/commit/20608a65548861a75a7ab2a63fc6c9f3fa8044b9))
* **card:** 🔘 增加按钮点击事件类型 ([ccd1d96](https://github.com/Aimerny/kook-go/commit/ccd1d967489dcbed823276f2e9d10ff0d24d466e))


### Bug Fixes

* 🐛 修复http_helper.go和message.go中的问题，优化代码逻辑 ([f14ec0c](https://github.com/Aimerny/kook-go/commit/f14ec0c8624e538a2f0be20fcb63495ae1847a80))
* **message:** 🐛 修复消息更新函数的返回结果处理 ([933bce9](https://github.com/Aimerny/kook-go/commit/933bce933537c6f623973ac8009334ccf074e4af))

## [1.2.2](https://github.com/Aimerny/kook-go/compare/v1.2.1...v1.2.2) (2024-07-08)


### Performance Improvements

* :zap: use jsoniter at whole project ([04d28d2](https://github.com/Aimerny/kook-go/commit/04d28d298f375cf42e75827e27548af43dbd3cc9))

## [1.2.1](https://github.com/Aimerny/kook-go/compare/v1.2.0...v1.2.1) (2024-06-01)


### Bug Fixes

* :bug: some card fields set omitempty ([a7bb94c](https://github.com/Aimerny/kook-go/commit/a7bb94c1cc0628b58f57761fe4b69d98615c7df2))

## [1.2.0](https://github.com/Aimerny/kook-go/compare/v1.1.1...v1.2.0) (2024-05-31)


### Features

* :sparkles: support card message action ([31b4f80](https://github.com/Aimerny/kook-go/commit/31b4f8099ab65849cf28f1d67dc654c8f527dd82))

## [1.1.1](https://github.com/Aimerny/kook-go/compare/v1.1.0...v1.1.1) (2024-05-31)


### Bug Fixes

* :bug: message req type use struct ([9f7b8a6](https://github.com/Aimerny/kook-go/commit/9f7b8a6f1ac0e2ccb861ae0d1921ccaf4c889aa2))

## [1.1.0](https://github.com/Aimerny/kook-go/compare/v1.0.0...v1.1.0) (2024-05-31)


### Features

* :fire: remove sdk config ([1fc0d1f](https://github.com/Aimerny/kook-go/commit/1fc0d1f1a988a761903c968ffd1309bb454f077a))


### Bug Fixes

* :bug: event type use struct instead ([bc42300](https://github.com/Aimerny/kook-go/commit/bc42300de7e0927f20728d1826838bad66f4423a))
