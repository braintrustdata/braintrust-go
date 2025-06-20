# Changelog

## 0.8.0 (2025-06-17)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/braintrustdata/braintrust-go/compare/v0.7.0...v0.8.0)

### ⚠ BREAKING CHANGES

* **client:** rename resp package
* **client:** improve core function names
* **client:** improve union variant names
* **client:** improve param subunions & deduplicate types

### Features

* **api:** manual updates ([#143](https://github.com/braintrustdata/braintrust-go/issues/143)) ([0ad70ec](https://github.com/braintrustdata/braintrust-go/commit/0ad70ec11f4b7d4c54d5273ae259c377cffc7819))
* **client:** add debug log helper ([3591364](https://github.com/braintrustdata/braintrust-go/commit/35913643a22c49b2f1adb855d1a920f72a92af7b))
* **client:** add escape hatch to omit required param fields ([#150](https://github.com/braintrustdata/braintrust-go/issues/150)) ([5058854](https://github.com/braintrustdata/braintrust-go/commit/5058854b03274d44549c17989dd6e951e43c911a))
* **client:** add support for endpoint-specific base URLs in python ([538ea47](https://github.com/braintrustdata/braintrust-go/commit/538ea47512fb9876a49df809d2a48c98c0577882))
* **client:** add support for reading base URL from environment variable ([febb457](https://github.com/braintrustdata/braintrust-go/commit/febb4578d89095ae52ff06104922b4c165a4a4d6))
* **client:** allow overriding unions ([e8e4fe3](https://github.com/braintrustdata/braintrust-go/commit/e8e4fe3be01361d1c18710661cc2d23b672fc47e))
* **client:** experimental support for unmarshalling into param structs ([af4406e](https://github.com/braintrustdata/braintrust-go/commit/af4406e238d0c280e95888e6c52b4b42245a8374))
* **client:** improve param subunions & deduplicate types ([3c0e907](https://github.com/braintrustdata/braintrust-go/commit/3c0e907a3ecf6b3bdc00ff33e603a6ed1f362edc))
* **client:** rename resp package ([7895a14](https://github.com/braintrustdata/braintrust-go/commit/7895a140f3131cb039601b411f11b2faa6140fe9))
* **client:** support custom http clients ([#152](https://github.com/braintrustdata/braintrust-go/issues/152)) ([bab0058](https://github.com/braintrustdata/braintrust-go/commit/bab00589900543a46510c691f32c064e18d8d785))
* **client:** support more time formats ([f7c7d39](https://github.com/braintrustdata/braintrust-go/commit/f7c7d39d1222649adaed7f6b8902cb394dd7f5f3))
* **client:** support param struct overrides ([#145](https://github.com/braintrustdata/braintrust-go/issues/145)) ([0baf2f6](https://github.com/braintrustdata/braintrust-go/commit/0baf2f64f4e8a809d26dcc7db7830f69eaf9ecc0))
* **client:** support unions in query and forms ([#148](https://github.com/braintrustdata/braintrust-go/issues/148)) ([c49965f](https://github.com/braintrustdata/braintrust-go/commit/c49965f1fdd38099fe927ab12f2eb5230af4b6f6))


### Bug Fixes

* **client:** cast to raw message when converting to params ([88b5d2f](https://github.com/braintrustdata/braintrust-go/commit/88b5d2fd0333090af6299046a038ee052321ac06))
* **client:** clean up reader resources ([ac1b426](https://github.com/braintrustdata/braintrust-go/commit/ac1b4267d58425c8460b2db0673ad388eacf9c6f))
* **client:** correctly set stream key for multipart ([9aeecf5](https://github.com/braintrustdata/braintrust-go/commit/9aeecf514d59861d97dfe82b4023b889e1d1b86e))
* **client:** correctly update body in WithJSONSet ([11a6d11](https://github.com/braintrustdata/braintrust-go/commit/11a6d11af861687140b494e354a4bb3319a3441b))
* **client:** don't panic on marshal with extra null field ([5546f82](https://github.com/braintrustdata/braintrust-go/commit/5546f827da3dfaf4abe55898b4b7ffa20b350aa9))
* **client:** improve core function names ([014a26a](https://github.com/braintrustdata/braintrust-go/commit/014a26a347b34df97e98e24d1aee9aa3b4633362))
* **client:** improve union variant names ([7fa4de7](https://github.com/braintrustdata/braintrust-go/commit/7fa4de7749644bae2127fceec8902ac9f378be8e))
* **client:** resolve issue with optional multipart files ([51a5e0b](https://github.com/braintrustdata/braintrust-go/commit/51a5e0b76a6225af03c58011e1593488e1125af9))
* **client:** return error on bad custom url instead of panic ([#146](https://github.com/braintrustdata/braintrust-go/issues/146)) ([b0d0fac](https://github.com/braintrustdata/braintrust-go/commit/b0d0fac421e4db8dc283fbb0465973581e398c49))
* **client:** support multipart encoding array formats ([#147](https://github.com/braintrustdata/braintrust-go/issues/147)) ([8eee61a](https://github.com/braintrustdata/braintrust-go/commit/8eee61a01d517a8ad10b1886d8f32a292b43e461))
* **client:** time format encoding fix ([a560d93](https://github.com/braintrustdata/braintrust-go/commit/a560d93727c2999ffa241139b6207353dd61c348))
* **client:** unmarshal responses properly ([9e2c24b](https://github.com/braintrustdata/braintrust-go/commit/9e2c24b0e8f1f4c7f926ee938b6509f8ce0a2d07))
* fix error ([f1257f0](https://github.com/braintrustdata/braintrust-go/commit/f1257f0e155aba6961fc60ab1e6589780721080f))
* handle empty bodies in WithJSONSet ([5003fac](https://github.com/braintrustdata/braintrust-go/commit/5003fac9849f41780ec7c3d53bb5541492989b5b))
* **pagination:** handle errors when applying options ([e4e55ff](https://github.com/braintrustdata/braintrust-go/commit/e4e55ffc43bf59fdda865d76b09830760ec926b5))
* **test:** return early after test failure ([#141](https://github.com/braintrustdata/braintrust-go/issues/141)) ([37c95eb](https://github.com/braintrustdata/braintrust-go/commit/37c95ebb24d4a6a415d882ae0e9ac741ba0085a4))


### Chores

* add request options to client tests ([#140](https://github.com/braintrustdata/braintrust-go/issues/140)) ([13d2104](https://github.com/braintrustdata/braintrust-go/commit/13d210467e84ee86d33e142a0cd844abe0885841))
* **ci:** add timeout thresholds for CI jobs ([b3345b6](https://github.com/braintrustdata/braintrust-go/commit/b3345b69e28c7ccc3e704fc71c48e4785ecfc0bf))
* **ci:** enable for pull requests ([e25c30c](https://github.com/braintrustdata/braintrust-go/commit/e25c30c7caee28aec6663154259870dd10206d5e))
* **ci:** only use depot for staging repos ([43b7288](https://github.com/braintrustdata/braintrust-go/commit/43b728868ca235f5f4acaa4b543df5aefdf652c5))
* **docs:** doc improvements ([#149](https://github.com/braintrustdata/braintrust-go/issues/149)) ([6bd9e2a](https://github.com/braintrustdata/braintrust-go/commit/6bd9e2aedca74e2c22ae7ddbb7955da1269bfce1))
* **docs:** document pre-request options ([2e2b1ba](https://github.com/braintrustdata/braintrust-go/commit/2e2b1ba5033c4aa2780ff07deed272734e5e3a0d))
* **docs:** grammar improvements ([16390fa](https://github.com/braintrustdata/braintrust-go/commit/16390faa8872b32927b1dff83f16fa3f8f60e80d))
* **docs:** improve security documentation ([#138](https://github.com/braintrustdata/braintrust-go/issues/138)) ([b8d2249](https://github.com/braintrustdata/braintrust-go/commit/b8d2249f32bd44f1463a03830906f6e53529ae92))
* **docs:** readme improvements ([#151](https://github.com/braintrustdata/braintrust-go/issues/151)) ([c7916fd](https://github.com/braintrustdata/braintrust-go/commit/c7916fda2d2a664db3381b7eb8a75ce230b65f76))
* **docs:** update file uploads in README ([#144](https://github.com/braintrustdata/braintrust-go/issues/144)) ([0970fb3](https://github.com/braintrustdata/braintrust-go/commit/0970fb3666bec584abe528f1833402157ac36f8a))
* **docs:** update respjson package name ([495df23](https://github.com/braintrustdata/braintrust-go/commit/495df236003b829ad8824a28d379ba30b95ac635))
* fix typos ([#142](https://github.com/braintrustdata/braintrust-go/issues/142)) ([6019d95](https://github.com/braintrustdata/braintrust-go/commit/6019d95f5d92016e03b4cbd20bb5d69602511fbf))
* improve devcontainer setup ([bc84a08](https://github.com/braintrustdata/braintrust-go/commit/bc84a085103697c13c910d0fda4901c42253e470))
* **internal:** codegen related update ([6d633ca](https://github.com/braintrustdata/braintrust-go/commit/6d633ca1720bee0651ad6f040ff25d38049d0492))
* **internal:** expand CI branch coverage ([068b8e7](https://github.com/braintrustdata/braintrust-go/commit/068b8e7f70aea5840664625fcaa5517eb760bb92))
* **internal:** reduce CI branch coverage ([4564b6a](https://github.com/braintrustdata/braintrust-go/commit/4564b6add83e1a7a73c4efdbe979b67dd245c449))
* make go mod tidy continue on error ([206aac6](https://github.com/braintrustdata/braintrust-go/commit/206aac6443db26806956de1bba6d38fc6d5e0e9b))
* **readme:** improve formatting ([3ecd3b8](https://github.com/braintrustdata/braintrust-go/commit/3ecd3b85fb3ff79fcc4b66074a1bd3d5fb61fc18))
* **utils:** add internal resp to param utility ([a681757](https://github.com/braintrustdata/braintrust-go/commit/a681757186db33e1e52c365e41bcffaf36531b86))


### Documentation

* update documentation links to be more uniform ([eefe3d3](https://github.com/braintrustdata/braintrust-go/commit/eefe3d37dfa4a2a5d55682c881ff052fc04c1b22))

## 0.7.0 (2025-03-14)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/braintrustdata/braintrust-go/compare/v0.6.0...v0.7.0)

### Features

* **api:** manual updates ([#135](https://github.com/braintrustdata/braintrust-go/issues/135)) ([b42ab00](https://github.com/braintrustdata/braintrust-go/commit/b42ab00d3ff3d2ce617c85e09dac892c6bf97082))

## 0.6.0 (2025-03-14)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/braintrustdata/braintrust-go/compare/v0.5.0...v0.6.0)

### Features

* add SKIP_BREW env var to ./scripts/bootstrap ([#127](https://github.com/braintrustdata/braintrust-go/issues/127)) ([d3d254e](https://github.com/braintrustdata/braintrust-go/commit/d3d254e3a205a4e08e650369635ab27047f6e505))
* **api:** manual updates ([#132](https://github.com/braintrustdata/braintrust-go/issues/132)) ([4156f1e](https://github.com/braintrustdata/braintrust-go/commit/4156f1eef8e761006513811a83468ad613edec6a))
* **api:** manual updates ([#133](https://github.com/braintrustdata/braintrust-go/issues/133)) ([846c8d4](https://github.com/braintrustdata/braintrust-go/commit/846c8d4069448d271c25ca8adbc610dc9cfb8ae4))
* **client:** accept RFC6838 JSON content types ([#128](https://github.com/braintrustdata/braintrust-go/issues/128)) ([03f3af7](https://github.com/braintrustdata/braintrust-go/commit/03f3af7eb119976bc20156e21af89c7a9f5eade6))
* **client:** allow custom baseurls without trailing slash ([#126](https://github.com/braintrustdata/braintrust-go/issues/126)) ([c175900](https://github.com/braintrustdata/braintrust-go/commit/c1759002dd0f06526907d198a10682dfe1508b04))
* **client:** improve default client options support ([#130](https://github.com/braintrustdata/braintrust-go/issues/130)) ([dbfcb1b](https://github.com/braintrustdata/braintrust-go/commit/dbfcb1b4ecccd9e6d20989ab51136259e219864d))
* **client:** send `X-Stainless-Timeout` header ([#118](https://github.com/braintrustdata/braintrust-go/issues/118)) ([80c1565](https://github.com/braintrustdata/braintrust-go/commit/80c156562081332a1c4a81310a84432beccd3ece))


### Bug Fixes

* **client:** don't truncate manually specified filenames ([#123](https://github.com/braintrustdata/braintrust-go/issues/123)) ([2e6f118](https://github.com/braintrustdata/braintrust-go/commit/2e6f1181beac39df27f9ddd3836613005e51b153))
* do not call path.Base on ContentType ([#122](https://github.com/braintrustdata/braintrust-go/issues/122)) ([c6cb6aa](https://github.com/braintrustdata/braintrust-go/commit/c6cb6aa7f982fd8cda9f60e4dd86f930575aef9f))
* fix apijson.Port for embedded structs ([#110](https://github.com/braintrustdata/braintrust-go/issues/110)) ([1791277](https://github.com/braintrustdata/braintrust-go/commit/1791277c48cd4d8eec13a962578405ad2668beca))
* fix apijson.Port for embedded structs ([#111](https://github.com/braintrustdata/braintrust-go/issues/111)) ([5301af0](https://github.com/braintrustdata/braintrust-go/commit/5301af03c238757d67d35dfed6eaf9b6da04e92e))
* fix early cancel when RequestTimeout is provided for streaming requests ([#120](https://github.com/braintrustdata/braintrust-go/issues/120)) ([5e636f7](https://github.com/braintrustdata/braintrust-go/commit/5e636f7952aec436910460a9a0deef6538e3a9fd))
* fix interface implementation stub names for unions ([#113](https://github.com/braintrustdata/braintrust-go/issues/113)) ([a89a6d5](https://github.com/braintrustdata/braintrust-go/commit/a89a6d50a799e54464b30fe6d31ba3920102c5e8))
* fix unicode encoding for json ([#116](https://github.com/braintrustdata/braintrust-go/issues/116)) ([4d28998](https://github.com/braintrustdata/braintrust-go/commit/4d2899885b9aea04b0d8255a9496e8c258ab02a9))


### Chores

* add UnionUnmarshaler for responses that are interfaces ([#119](https://github.com/braintrustdata/braintrust-go/issues/119)) ([ac3c813](https://github.com/braintrustdata/braintrust-go/commit/ac3c8132b8392a91e044cdfa01e78749487a0335))
* **client:** remove unnecessary `nullable` struct tags ([#107](https://github.com/braintrustdata/braintrust-go/issues/107)) ([67da890](https://github.com/braintrustdata/braintrust-go/commit/67da89065e7b4afda857129e6c6903402bb930de))
* **internal:** codegen related update ([#105](https://github.com/braintrustdata/braintrust-go/issues/105)) ([953511f](https://github.com/braintrustdata/braintrust-go/commit/953511f3de68a0829dae360a3f8044b80b1a08fe))
* **internal:** codegen related update ([#106](https://github.com/braintrustdata/braintrust-go/issues/106)) ([386543b](https://github.com/braintrustdata/braintrust-go/commit/386543bd410409b5654cf216d632a4371eb854f2))
* **internal:** codegen related update ([#108](https://github.com/braintrustdata/braintrust-go/issues/108)) ([8b0908f](https://github.com/braintrustdata/braintrust-go/commit/8b0908fa059ee934bdb7f22bd9c9d6f28baacfbf))
* **internal:** codegen related update ([#109](https://github.com/braintrustdata/braintrust-go/issues/109)) ([72bd256](https://github.com/braintrustdata/braintrust-go/commit/72bd25617efca17a6d62ac3678414c52989f3ebc))
* **internal:** codegen related update ([#112](https://github.com/braintrustdata/braintrust-go/issues/112)) ([50142f6](https://github.com/braintrustdata/braintrust-go/commit/50142f61a74e4f245b4c28d84b589e47c88887e1))
* **internal:** codegen related update ([#115](https://github.com/braintrustdata/braintrust-go/issues/115)) ([dec22c9](https://github.com/braintrustdata/braintrust-go/commit/dec22c9b9e21ed3d799d9c45f6f40ef9381fee28))
* **internal:** fix devcontainers setup ([#124](https://github.com/braintrustdata/braintrust-go/issues/124)) ([19dcac6](https://github.com/braintrustdata/braintrust-go/commit/19dcac62f53372200925b74fd97e3b2b63d90bbd))
* **internal:** remove extra empty newlines ([#131](https://github.com/braintrustdata/braintrust-go/issues/131)) ([3e8a639](https://github.com/braintrustdata/braintrust-go/commit/3e8a639ea5d703e71c023d44c1f64033e9b592c6))
* **internal:** version bump ([#103](https://github.com/braintrustdata/braintrust-go/issues/103)) ([b7b6ac8](https://github.com/braintrustdata/braintrust-go/commit/b7b6ac85800f427e9095ae80203bcb392b36a442))
* minor change to tests ([#121](https://github.com/braintrustdata/braintrust-go/issues/121)) ([396d6c9](https://github.com/braintrustdata/braintrust-go/commit/396d6c9ef39c5916044da3fcaf663fd36d4371d9))
* refactor client tests ([#114](https://github.com/braintrustdata/braintrust-go/issues/114)) ([eba2c08](https://github.com/braintrustdata/braintrust-go/commit/eba2c08ed874f6cc061d945f48fd08d259f35a95))


### Documentation

* document raw responses ([#117](https://github.com/braintrustdata/braintrust-go/issues/117)) ([e7502ce](https://github.com/braintrustdata/braintrust-go/commit/e7502ce0ab2348c3cbf84145c062285a3b026b29))
* update URLs from stainlessapi.com to stainless.com ([#125](https://github.com/braintrustdata/braintrust-go/issues/125)) ([8c12a9d](https://github.com/braintrustdata/braintrust-go/commit/8c12a9d0c5ac918e171d236291e3bf4ab43a59c8))


### Refactors

* tidy up dependencies ([#129](https://github.com/braintrustdata/braintrust-go/issues/129)) ([f24edf4](https://github.com/braintrustdata/braintrust-go/commit/f24edf488db2320e948aff01a4eafce81c220de9))

## 0.5.0 (2024-11-20)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/braintrustdata/braintrust-go/compare/v0.4.0...v0.5.0)

### Features

* **api:** manual updates ([#100](https://github.com/braintrustdata/braintrust-go/issues/100)) ([7cc183a](https://github.com/braintrustdata/braintrust-go/commit/7cc183acd30fba1bcdff0f6e1357492d38183bc3))
* **api:** manual updates ([#101](https://github.com/braintrustdata/braintrust-go/issues/101)) ([3c9419c](https://github.com/braintrustdata/braintrust-go/commit/3c9419c3b959a101a83ebdd255a10fb1d78552d5))
* **api:** manual updates ([#92](https://github.com/braintrustdata/braintrust-go/issues/92)) ([e7dd050](https://github.com/braintrustdata/braintrust-go/commit/e7dd050272cfede0031c11dfcf8d18d0cecdbd0e))
* **api:** manual updates ([#93](https://github.com/braintrustdata/braintrust-go/issues/93)) ([fe6e368](https://github.com/braintrustdata/braintrust-go/commit/fe6e36847f7e12d16fc0ba956de19b34f15a726e))
* **api:** manual updates ([#94](https://github.com/braintrustdata/braintrust-go/issues/94)) ([b88818e](https://github.com/braintrustdata/braintrust-go/commit/b88818edf9a44eeae32c368242fc9dd1b88cd179))
* **api:** manual updates ([#96](https://github.com/braintrustdata/braintrust-go/issues/96)) ([be7e542](https://github.com/braintrustdata/braintrust-go/commit/be7e542d6512f5d83c9cb5d4d48ebc91b95874fd))
* **api:** manual updates ([#97](https://github.com/braintrustdata/braintrust-go/issues/97)) ([a954b64](https://github.com/braintrustdata/braintrust-go/commit/a954b641a274bfdec74acc36183c6ed440da6f7b))
* **api:** manual updates ([#98](https://github.com/braintrustdata/braintrust-go/issues/98)) ([40d7dc7](https://github.com/braintrustdata/braintrust-go/commit/40d7dc715b459f8e78aa141fe22879f6a473b0c4))
* **api:** manual updates ([#99](https://github.com/braintrustdata/braintrust-go/issues/99)) ([78c7f99](https://github.com/braintrustdata/braintrust-go/commit/78c7f99b766efbf8a342b98a6c84da38d2d22866))


### Chores

* rebuild project due to codegen change ([#87](https://github.com/braintrustdata/braintrust-go/issues/87)) ([8bebfe9](https://github.com/braintrustdata/braintrust-go/commit/8bebfe9470a2cf8afec8b24370b53163e985ce69))
* rebuild project due to codegen change ([#89](https://github.com/braintrustdata/braintrust-go/issues/89)) ([aa4fc32](https://github.com/braintrustdata/braintrust-go/commit/aa4fc32785a306c6e57a8e6a4d2a857b36905cb5))
* rebuild project due to codegen change ([#90](https://github.com/braintrustdata/braintrust-go/issues/90)) ([7118289](https://github.com/braintrustdata/braintrust-go/commit/7118289f424023992d87e42893675ff3fa3d4d81))
* rebuild project due to codegen change ([#91](https://github.com/braintrustdata/braintrust-go/issues/91)) ([2053c0a](https://github.com/braintrustdata/braintrust-go/commit/2053c0ad42dbed139cf00ca10fd38e57174bec6c))
* rebuild project due to codegen change ([#95](https://github.com/braintrustdata/braintrust-go/issues/95)) ([a0b6db1](https://github.com/braintrustdata/braintrust-go/commit/a0b6db132d720ed9942ffe9590905c28aca858f1))

## 0.4.0 (2024-10-28)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/braintrustdata/braintrust-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** api update ([#84](https://github.com/braintrustdata/braintrust-go/issues/84)) ([df1045e](https://github.com/braintrustdata/braintrust-go/commit/df1045efbcc599b0ae883736f2ea15f1a6aa8af4))
* **api:** deduplication ([#79](https://github.com/braintrustdata/braintrust-go/issues/79)) ([5686ac7](https://github.com/braintrustdata/braintrust-go/commit/5686ac7a65d9c0e8bb72fcebfbc9aac3452c1d2e))
* **api:** manual updates ([#80](https://github.com/braintrustdata/braintrust-go/issues/80)) ([e3165d0](https://github.com/braintrustdata/braintrust-go/commit/e3165d0c7b0c670a4f29c944174f87c9e47d4164))
* **api:** manual updates ([#82](https://github.com/braintrustdata/braintrust-go/issues/82)) ([ba65282](https://github.com/braintrustdata/braintrust-go/commit/ba65282695f9e8a52aef0cb10c146929e3f973bd))
* **api:** manual updates ([#83](https://github.com/braintrustdata/braintrust-go/issues/83)) ([10663e1](https://github.com/braintrustdata/braintrust-go/commit/10663e14b30b7ffaedb366ffe743579cb198f1db))
* **api:** update via SDK Studio ([#64](https://github.com/braintrustdata/braintrust-go/issues/64)) ([5d94a4a](https://github.com/braintrustdata/braintrust-go/commit/5d94a4a13a1609d12f6d22dc1ed1822153fbaf55))
* **api:** update via SDK Studio ([#67](https://github.com/braintrustdata/braintrust-go/issues/67)) ([8690ed8](https://github.com/braintrustdata/braintrust-go/commit/8690ed8502bdbaa75a564da00c9c33fb94a02b9e))
* **api:** update via SDK Studio ([#68](https://github.com/braintrustdata/braintrust-go/issues/68)) ([2d4f632](https://github.com/braintrustdata/braintrust-go/commit/2d4f6323a6394d5416b8f02978c8b9354a53728b))
* **api:** update via SDK Studio ([#69](https://github.com/braintrustdata/braintrust-go/issues/69)) ([a0f9c5b](https://github.com/braintrustdata/braintrust-go/commit/a0f9c5b0d822a67282fa874021f9dae9f92a1baf))
* **api:** update via SDK Studio ([#70](https://github.com/braintrustdata/braintrust-go/issues/70)) ([3aef711](https://github.com/braintrustdata/braintrust-go/commit/3aef7112b35e37d846426c6772ddb45d68abeac8))
* **api:** update via SDK Studio ([#71](https://github.com/braintrustdata/braintrust-go/issues/71)) ([12ff266](https://github.com/braintrustdata/braintrust-go/commit/12ff266ee7b8527a72d07ceebbc1b4f8be1c5d1c))
* **api:** update via SDK Studio ([#72](https://github.com/braintrustdata/braintrust-go/issues/72)) ([374f708](https://github.com/braintrustdata/braintrust-go/commit/374f7088260e53044dc1a72bcf38f955792f82f8))
* **api:** update via SDK Studio ([#73](https://github.com/braintrustdata/braintrust-go/issues/73)) ([2df3cf4](https://github.com/braintrustdata/braintrust-go/commit/2df3cf43958dbebd6504821150d94c1f97c2f32f))
* **api:** update via SDK Studio ([#74](https://github.com/braintrustdata/braintrust-go/issues/74)) ([900928b](https://github.com/braintrustdata/braintrust-go/commit/900928bf8480236b28113e712abb6e0dd2c0a36e))
* **api:** update via SDK Studio ([#75](https://github.com/braintrustdata/braintrust-go/issues/75)) ([9b6d29e](https://github.com/braintrustdata/braintrust-go/commit/9b6d29eed55d2aba4a283711b308e27508da3c5e))
* **api:** update via SDK Studio ([#76](https://github.com/braintrustdata/braintrust-go/issues/76)) ([55bf0e9](https://github.com/braintrustdata/braintrust-go/commit/55bf0e9df88c59bc10f4b4dfee549c7307bea81a))


### Bug Fixes

* **api:** fix go build ([#78](https://github.com/braintrustdata/braintrust-go/issues/78)) ([0836190](https://github.com/braintrustdata/braintrust-go/commit/0836190bbc023d57751080d29b5d24cd9d54e6ab))
* **api:** fix go unions ([#81](https://github.com/braintrustdata/braintrust-go/issues/81)) ([3407648](https://github.com/braintrustdata/braintrust-go/commit/3407648208167e78fe9e5179e551dd0626371c9c))


### Chores

* **api:** manual updates ([#85](https://github.com/braintrustdata/braintrust-go/issues/85)) ([277aa99](https://github.com/braintrustdata/braintrust-go/commit/277aa99a39816b54882e0f6fffb30ceab4f274c4))
* **internal:** codegen related update ([#65](https://github.com/braintrustdata/braintrust-go/issues/65)) ([9a1b423](https://github.com/braintrustdata/braintrust-go/commit/9a1b4236981a93ef2ffa28441b7bebc2ca8cabc2))
* **internal:** codegen related update ([#77](https://github.com/braintrustdata/braintrust-go/issues/77)) ([371af04](https://github.com/braintrustdata/braintrust-go/commit/371af04dce69372c9583f6aa0d3f88924189d14f))

## 0.3.0 (2024-08-09)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/braintrustdata/braintrust-go/compare/v0.2.0...v0.3.0)

### Features

* **api:** update via SDK Studio ([#60](https://github.com/braintrustdata/braintrust-go/issues/60)) ([0428527](https://github.com/braintrustdata/braintrust-go/commit/04285278647f83c4658d1afbf4dbd196331ef3e7))


### Bug Fixes

* deserialization of struct unions that implement json.Unmarshaler ([#62](https://github.com/braintrustdata/braintrust-go/issues/62)) ([9abee30](https://github.com/braintrustdata/braintrust-go/commit/9abee30e0c3db70118e891f17f427d03bbe045a7))

## 0.2.0 (2024-08-09)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/braintrustdata/braintrust-go/compare/v0.1.0...v0.2.0)

### Features

* add model ProjectScoreCategory ([#53](https://github.com/braintrustdata/braintrust-go/issues/53)) ([df3257b](https://github.com/braintrustdata/braintrust-go/commit/df3257b4182d17511efdf2fc8d15ea4e84b8d5f6))
* **api:** manual updates ([#54](https://github.com/braintrustdata/braintrust-go/issues/54)) ([805aab0](https://github.com/braintrustdata/braintrust-go/commit/805aab0e2c9626bd51e131bf3bf5fbe06186adde))
* **api:** manual updates ([#56](https://github.com/braintrustdata/braintrust-go/issues/56)) ([03c070d](https://github.com/braintrustdata/braintrust-go/commit/03c070d08873ee6af05bc2346c01caebdf982036))
* **api:** manual updates ([#57](https://github.com/braintrustdata/braintrust-go/issues/57)) ([39346f6](https://github.com/braintrustdata/braintrust-go/commit/39346f6a70933ceb8e0769b2f45282a48a555762))
* **api:** update via SDK Studio ([#10](https://github.com/braintrustdata/braintrust-go/issues/10)) ([f0a86bd](https://github.com/braintrustdata/braintrust-go/commit/f0a86bdf968ed2bb74b914775264415e9ac9b7da))
* **api:** update via SDK Studio ([#11](https://github.com/braintrustdata/braintrust-go/issues/11)) ([177ba23](https://github.com/braintrustdata/braintrust-go/commit/177ba23b86917fa9c053ea5e6de9366fbd8fcac7))
* **api:** update via SDK Studio ([#12](https://github.com/braintrustdata/braintrust-go/issues/12)) ([4327e1f](https://github.com/braintrustdata/braintrust-go/commit/4327e1f8e70cdd2623b8ffd4b6b7f862e2aa84fa))
* **api:** update via SDK Studio ([#13](https://github.com/braintrustdata/braintrust-go/issues/13)) ([8a9748d](https://github.com/braintrustdata/braintrust-go/commit/8a9748d48c708c35caea0b9d52e756c10e3f5db3))
* **api:** update via SDK Studio ([#14](https://github.com/braintrustdata/braintrust-go/issues/14)) ([2c0e8ce](https://github.com/braintrustdata/braintrust-go/commit/2c0e8cef0d7e0458367b994b78257c7041d00b10))
* **api:** update via SDK Studio ([#15](https://github.com/braintrustdata/braintrust-go/issues/15)) ([a8ff61d](https://github.com/braintrustdata/braintrust-go/commit/a8ff61d0c6564fe990a092c12a41bc3aba927d59))
* **api:** update via SDK Studio ([#16](https://github.com/braintrustdata/braintrust-go/issues/16)) ([60895f1](https://github.com/braintrustdata/braintrust-go/commit/60895f14774b0892b517d28316a036aadb40823e))
* **api:** update via SDK Studio ([#17](https://github.com/braintrustdata/braintrust-go/issues/17)) ([c69d558](https://github.com/braintrustdata/braintrust-go/commit/c69d55859f999c8e5e1b84289a9540b7d13a54df))
* **api:** update via SDK Studio ([#18](https://github.com/braintrustdata/braintrust-go/issues/18)) ([113fd22](https://github.com/braintrustdata/braintrust-go/commit/113fd221b58f6d9d225b29f9668563c34d076a61))
* **api:** update via SDK Studio ([#19](https://github.com/braintrustdata/braintrust-go/issues/19)) ([1da3d3c](https://github.com/braintrustdata/braintrust-go/commit/1da3d3c329272988ba847ecb716d6184d17a91c6))
* **api:** update via SDK Studio ([#20](https://github.com/braintrustdata/braintrust-go/issues/20)) ([fd8c49b](https://github.com/braintrustdata/braintrust-go/commit/fd8c49bd693c717d176e253f37c7a5555760b683))
* **api:** update via SDK Studio ([#22](https://github.com/braintrustdata/braintrust-go/issues/22)) ([2dc1dba](https://github.com/braintrustdata/braintrust-go/commit/2dc1dba976adf081f78a3c6494871a502fd3e0a1))
* **api:** update via SDK Studio ([#28](https://github.com/braintrustdata/braintrust-go/issues/28)) ([089d65a](https://github.com/braintrustdata/braintrust-go/commit/089d65a839ffc7144b627339fcd1ca89ccffc793))
* **api:** update via SDK Studio ([#29](https://github.com/braintrustdata/braintrust-go/issues/29)) ([2baf1a1](https://github.com/braintrustdata/braintrust-go/commit/2baf1a1b5666cdd92cd38ccf96e24d233499f779))
* **api:** update via SDK Studio ([#33](https://github.com/braintrustdata/braintrust-go/issues/33)) ([e8f4aa7](https://github.com/braintrustdata/braintrust-go/commit/e8f4aa7a72d8026c2f0246c1b716bcfab047cd20))
* **api:** update via SDK Studio ([#34](https://github.com/braintrustdata/braintrust-go/issues/34)) ([4a586a7](https://github.com/braintrustdata/braintrust-go/commit/4a586a7a02f872a3c8b522e6bbdaed7111a58570))
* **api:** update via SDK Studio ([#35](https://github.com/braintrustdata/braintrust-go/issues/35)) ([1d783d6](https://github.com/braintrustdata/braintrust-go/commit/1d783d640927e87edef070e7632c758ad1fef4d4))
* **api:** update via SDK Studio ([#36](https://github.com/braintrustdata/braintrust-go/issues/36)) ([737dd53](https://github.com/braintrustdata/braintrust-go/commit/737dd53e4f38e9288244a4ffd6d5da85c57968de))
* **api:** update via SDK Studio ([#37](https://github.com/braintrustdata/braintrust-go/issues/37)) ([c113455](https://github.com/braintrustdata/braintrust-go/commit/c1134557ab766db92cd44164b7e8e979fa8bb677))
* **api:** update via SDK Studio ([#38](https://github.com/braintrustdata/braintrust-go/issues/38)) ([e68455f](https://github.com/braintrustdata/braintrust-go/commit/e68455f72e58a0d0f7159bdc5436cb983a2fcdc8))
* **api:** update via SDK Studio ([#39](https://github.com/braintrustdata/braintrust-go/issues/39)) ([9a9df62](https://github.com/braintrustdata/braintrust-go/commit/9a9df62edcdb41586590ef1160df31b3e3059d5c))
* **api:** update via SDK Studio ([#40](https://github.com/braintrustdata/braintrust-go/issues/40)) ([ded22c9](https://github.com/braintrustdata/braintrust-go/commit/ded22c934fbdaaf686ccdf62c0cbe30cc39e412d))
* **api:** update via SDK Studio ([#41](https://github.com/braintrustdata/braintrust-go/issues/41)) ([8839616](https://github.com/braintrustdata/braintrust-go/commit/8839616dbe0bb2eed4f12d82cec5373b5eae9e3b))
* **api:** update via SDK Studio ([#42](https://github.com/braintrustdata/braintrust-go/issues/42)) ([4ea34b0](https://github.com/braintrustdata/braintrust-go/commit/4ea34b06f155de39062d714a53e616457044c323))
* **api:** update via SDK Studio ([#43](https://github.com/braintrustdata/braintrust-go/issues/43)) ([40703d6](https://github.com/braintrustdata/braintrust-go/commit/40703d611cf092c1ead533beb2ac3d13df582dc6))
* **api:** update via SDK Studio ([#44](https://github.com/braintrustdata/braintrust-go/issues/44)) ([d28a4f5](https://github.com/braintrustdata/braintrust-go/commit/d28a4f55392d5ab58f275ba2485c31f0d4de7b31))
* **api:** update via SDK Studio ([#45](https://github.com/braintrustdata/braintrust-go/issues/45)) ([9e62cf7](https://github.com/braintrustdata/braintrust-go/commit/9e62cf7f745eb968939dc9d58ebf8b8ef78740bf))
* **api:** update via SDK Studio ([#46](https://github.com/braintrustdata/braintrust-go/issues/46)) ([351a97d](https://github.com/braintrustdata/braintrust-go/commit/351a97d86062a511e6dbcf9b6594455d81215025))
* **api:** update via SDK Studio ([#47](https://github.com/braintrustdata/braintrust-go/issues/47)) ([8f90f7c](https://github.com/braintrustdata/braintrust-go/commit/8f90f7cc150cdeb82ef7b9dde5c363e70ba3878c))
* **api:** update via SDK Studio ([#48](https://github.com/braintrustdata/braintrust-go/issues/48)) ([2ea300c](https://github.com/braintrustdata/braintrust-go/commit/2ea300c21179f5d9e3260c0fdf707b7e24e49920))
* **api:** update via SDK Studio ([#50](https://github.com/braintrustdata/braintrust-go/issues/50)) ([b1c2e3b](https://github.com/braintrustdata/braintrust-go/commit/b1c2e3b440232447ad6b3f1772ce01eb029fc27f))
* **api:** update via SDK Studio ([#52](https://github.com/braintrustdata/braintrust-go/issues/52)) ([0d622d3](https://github.com/braintrustdata/braintrust-go/commit/0d622d365c4377993c576b74a86b6d921188825c))
* **api:** update via SDK Studio ([#6](https://github.com/braintrustdata/braintrust-go/issues/6)) ([26d1406](https://github.com/braintrustdata/braintrust-go/commit/26d140665507b0688bb3be6f7cd26ebdbd376b0b))
* **api:** update via SDK Studio ([#8](https://github.com/braintrustdata/braintrust-go/issues/8)) ([9d22284](https://github.com/braintrustdata/braintrust-go/commit/9d22284f9d6a3645ee74a090fc296f6b603bdbc9))
* **api:** update via SDK Studio ([#9](https://github.com/braintrustdata/braintrust-go/issues/9)) ([4bd61e5](https://github.com/braintrustdata/braintrust-go/commit/4bd61e515179eee507ec848b0246fccf70943f44))


### Chores

* **ci:** bump prism mock server version ([#58](https://github.com/braintrustdata/braintrust-go/issues/58)) ([649c1a9](https://github.com/braintrustdata/braintrust-go/commit/649c1a98283165a0352c337bc84e92c9fa2bc322))
* **ci:** limit release doctor target branches ([#25](https://github.com/braintrustdata/braintrust-go/issues/25)) ([093856a](https://github.com/braintrustdata/braintrust-go/commit/093856a90d39f6b67f3cfbef86cf3b9fe2d688d5))
* **ci:** remove unused release doctor ([#26](https://github.com/braintrustdata/braintrust-go/issues/26)) ([ddc0d6a](https://github.com/braintrustdata/braintrust-go/commit/ddc0d6a81e389eb2cbdc453ae3343273d9e8e50b))
* **internal:** codegen related update ([#23](https://github.com/braintrustdata/braintrust-go/issues/23)) ([c542303](https://github.com/braintrustdata/braintrust-go/commit/c5423030ddd8907243a588a5e9ceeb3d8f4aee0d))
* **internal:** codegen related update ([#30](https://github.com/braintrustdata/braintrust-go/issues/30)) ([07677c4](https://github.com/braintrustdata/braintrust-go/commit/07677c43a9ba523cc3783d88369e63a52c9dc47b))
* **internal:** codegen related update ([#31](https://github.com/braintrustdata/braintrust-go/issues/31)) ([3bc15d3](https://github.com/braintrustdata/braintrust-go/commit/3bc15d3afddccd000ce6f4176e11846d6f71d400))
* **internal:** codegen related update ([#32](https://github.com/braintrustdata/braintrust-go/issues/32)) ([810e0fb](https://github.com/braintrustdata/braintrust-go/commit/810e0fb4bfde625c29079eee59acbf693e49ff6d))
* **internal:** codegen related update ([#49](https://github.com/braintrustdata/braintrust-go/issues/49)) ([dbb7f18](https://github.com/braintrustdata/braintrust-go/commit/dbb7f1842930799fc099dff46a1e3906083f1d0e))
* **internal:** codegen related update ([#51](https://github.com/braintrustdata/braintrust-go/issues/51)) ([c5ab4c7](https://github.com/braintrustdata/braintrust-go/commit/c5ab4c77e753138783757f2328db883d207c7075))
* **internal:** codegen related update ([#55](https://github.com/braintrustdata/braintrust-go/issues/55)) ([dace42a](https://github.com/braintrustdata/braintrust-go/commit/dace42a333b9cb4a7a799162eb5faada212d805d))
* **internal:** minor changes to tests ([#24](https://github.com/braintrustdata/braintrust-go/issues/24)) ([f5654d1](https://github.com/braintrustdata/braintrust-go/commit/f5654d1ea40735599feba445f0a116664ad1dba4))
* **tests:** update prism version ([#27](https://github.com/braintrustdata/braintrust-go/issues/27)) ([2b89c51](https://github.com/braintrustdata/braintrust-go/commit/2b89c51402ad3e7b46f35e515b3f57151e1e3130))
* update SDK settings ([#21](https://github.com/braintrustdata/braintrust-go/issues/21)) ([6ff7a70](https://github.com/braintrustdata/braintrust-go/commit/6ff7a7010bcde3ed20fc772c7c20f46f4818dd3c))

## 0.1.0 (2024-02-02)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/braintrustdata/braintrust-go/compare/v0.0.1...v0.1.0)

### Features

* Add initial Stainless SDK ([e0c9703](https://github.com/braintrustdata/braintrust-go/commit/e0c9703f875a8dfc061646b4989b248345f156fa))
* create default branch ([8db8d16](https://github.com/braintrustdata/braintrust-go/commit/8db8d165a825eee0342dffea7f7a62a798699884))
* OpenAPI spec update ([#1](https://github.com/braintrustdata/braintrust-go/issues/1)) ([6c2dcdb](https://github.com/braintrustdata/braintrust-go/commit/6c2dcdb3b50001e3fc3d9cafa223ba840e170b1b))
* OpenAPI spec update ([#3](https://github.com/braintrustdata/braintrust-go/issues/3)) ([fa58787](https://github.com/braintrustdata/braintrust-go/commit/fa58787337bafb43c2012f0ae3b9c04e5c5c72be))
* OpenAPI spec update ([#4](https://github.com/braintrustdata/braintrust-go/issues/4)) ([e660719](https://github.com/braintrustdata/braintrust-go/commit/e660719cc1e427410122e3062daca8904dcdfa3c))
* OpenAPI spec update ([#5](https://github.com/braintrustdata/braintrust-go/issues/5)) ([bb1c4b7](https://github.com/braintrustdata/braintrust-go/commit/bb1c4b7bf0bb5bdf280121a61814dc4e71719481))
