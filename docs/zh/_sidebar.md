- [首页](/zh/README.md)
- [快速开始](/zh/start/)
    - [使用State API](zh/start/state/start.md)
    - 使用Configuration API
        - [使用Apollo配置中心](zh/start/configuration/start-apollo.md)
        - [使用Etcd配置中心](zh/start/configuration/start.md)
    - [使用Pub/Sub API](zh/start/pubsub/start.md)
    - [使用分布式锁 API](zh/start/lock/start.md)
    - [使用Sequencer API生成分布式自增id](zh/start/sequencer/start.md)
    - 使用 Secret API
        - [查询 secrets](zh/start/secret/start.md)
        - [在组件配置中引用 secret](zh/start/secret/secret_ref.md)
    - 进行RPC调用
        - [Hello World](zh/start/rpc/helloworld.md)
        - [Dubbo JSON RPC](zh/start/rpc/dubbo_json_rpc.md)
    - 使用File API
        - [基于Minio](zh/start/file/minio.md)
    - [使用 OSS API](zh/start/oss/oss.md)
    <!--quickstart_generator-->
    - [(建设中)使用 phone API](zh/start/phone/start) 
    - [(建设中)使用 email API](zh/start/email/start) 
    - [使用 lifecycle API](zh/start/lifecycle/start) 
    - [API插件：注册您自己的API](zh/start/api_plugin/helloworld.md)
    - 作为 Istio 的数据面
        - [集成 Istio 1.10.6 演示](zh/start/istio/)
        - [集成 Istio 1.5.x 演示](zh/start/istio/start.md)
    - 在四层网络进行流量干预
        - [Dump TCP 流量](zh/start/network_filter/tcpcopy.md)
    - 在七层网络进行流量干预
        - [方法级别限流](zh/start/stream_filter/flow_control.md)
    - [健康检查、查询运行时元数据](zh/start/actuator/start.md)
    - 可观测性
        - [Trace, Metrics](zh/start/trace/trace.md)
        - [Trace 接入 Skywalking](zh/start/trace/skywalking.md)
        - [Trace 接入 Zipkin](zh/start/trace/zipkin.md)
        - [Trace 接入 Jaeger](zh/start/trace/jaeger.md)
        - [Metrics 接入 Prometheus](zh/start/trace/prometheus.md)
    - [将业务逻辑通过 WASM 下沉进sidecar](zh/start/wasm/start.md)
    - [基于 WASM 跟 Runtime 实现的 Faas 模型](zh/start/faas/start.md)
- [线上实验室](zh/start/lab.md)
- [用户手册](zh/building_blocks/)
    - 功能介绍
        - [File API](zh/building_blocks/file/file.md)
        - [Actuator API](zh/building_blocks/actuator/actuator.md)
        - [State API](zh/building_blocks/state/reference.md)
        - [Sequencer API](zh/building_blocks/sequencer/reference.md)
        - [Distributed Lock API](zh/building_blocks/lock/reference.md)
        - [Pub/Sub API](zh/building_blocks/pubsub/reference.md)
        - [RPC API](zh/building_blocks/rpc/reference.md)
        - [Configuration API](zh/building_blocks/configuration/reference.md)
        - 可扩展性
            - [API插件](zh/design/api_plugin/design.md)
    - [gRPC API 接口文档](zh/api_reference/README.md)
    - SDK文档
        - [java sdk](https://github.com/layotto/java-sdk)
        - [.net sdk](https://github.com/layotto/dotnet-sdk)
        - [js sdk](https://github.com/layotto/js-sdk)
        - [go sdk](zh/sdk_reference/go/start.md)
- 运维手册
    - [如何配置 Layotto](zh/configuration/)
        - [Layotto 配置文件介绍](zh/configuration/overview.md)
        - [组件配置说明](zh/component_specs/overview.md)
            - [State](zh/component_specs/state/common.md)
                - [Redis](zh/component_specs/state/redis.md)
                - [其他组件](zh/component_specs/state/others.md)
            - [Pub/Sub](zh/component_specs/pubsub/common.md)
                - [Redis](zh/component_specs/pubsub/redis.md)
                - [其他组件](zh/component_specs/pubsub/others.md)
            - [Distributed Lock](zh/component_specs/lock/common.md)
                - [Redis](zh/component_specs/lock/redis.md)
                - [Etcd](zh/component_specs/lock/etcd.md)
                - [Zookeeper](zh/component_specs/lock/zookeeper.md)
                - [Consul](zh/component_specs/lock/consul.md)
                - [MongoDB](zh/component_specs/lock/mongo.md)
            - Configuration
                - [Etcd](zh/component_specs/configuration/etcd.md)
                - [Apollo](zh/component_specs/configuration/apollo.md)
            - [File](zh/component_specs/file/common.md)
                - [OSS](zh/component_specs/file/oss.md)
            - [Sequencer](zh/component_specs/sequencer/common.md)
                - [Etcd](zh/component_specs/sequencer/etcd.md)
                - [Redis](zh/component_specs/sequencer/redis.md)
                - [Zookeeper](zh/component_specs/sequencer/zookeeper.md)
                - [MongoDB](zh/component_specs/sequencer/mongo.md)
                - [Mysql](zh/component_specs/sequencer/mysql.md)
            - [Secret Store](zh/component_specs/secret/common.md)  
            - [自定义组件](zh/component_specs/custom/common.md)
    - [如何部署、升级 Layotto](zh/operation/)
    - [如何本地开发、本地调试](zh/operation/local.md)
- 设计文档
    - [动态配置下发、组件热重载](zh/design/lifecycle/apply_configuration.md)
    - [Actuator设计文档](zh/design/actuator/actuator-design-doc.md)
    - [gRPC框架设计文档](zh/design/actuator/grpc-design-doc.md)
    - [Configuration API with Apollo](zh/design/configuration/configuration-api-with-apollo.md)
    - [Pub/Sub API以及与dapr component的兼容性](zh/design/pubsub/pubsub-api-and-compability-with-dapr-component.md)
    - [RPC设计文档](zh/design/rpc/rpc设计文档.md)
    - [分布式锁API设计文档](zh/design/lock/lock-api-design.md)
    - [Sequencer API设计文档](zh/design/sequencer/design.md)
    - [File API设计文档](zh/design/file/file-design.md)
    - [FaaS 设计文档](zh/design/faas/faas-poc-design.md)
    - [API插件](zh/design/api_plugin/design.md)
    - [支持Dapr API](zh/design/api_plugin/dapr_api.md)
    - [OSS API设计文档](zh/design/oss/oss-api-design.md)
- 贡献指南
    - [Layotto 贡献指南](zh/development/CONTRIBUTING.md)
    - [新手攻略：从零开始成为 Layotto 贡献者](zh/development/start-from-zero.md)
    - 想要贡献文档?
        - [文档贡献指南](zh/development/contributing-doc.md)
        - [使用工具自动测试 Quickstart 文档](zh/development/test-quickstart.md)
    - [想要开发新的组件?](zh/development/developing-component.md)
    - 想要修改proto文件或API定义？
        - [如何基于proto文件生成代码、接口文档](zh/api_reference/how_to_generate_api_doc.md)
        - [proto文件注释规范](zh/api_reference/comment_spec_of_proto.md)
        - [新增API时的开发规范](zh/development/developing-api.md)
    - [Layotto 四大 Github Workflows 说明](zh/development/github-workflows.md)
    - [Layotto 命令行工具指南](zh/development/commands.md)
    - 如何给 issue 打 label
        - [新手任务 (good first issue) 的 label 规范](zh/development/label-spec.md)
    - [发布手册](zh/development/release-guide.md)
    - [待解决的问题](zh/development/problems-to-solve.md)
- 社区
    - [社区会议](zh/community/meeting.md)
    - [SOFAStack & MOSN 社区角色说明](zh/community/governance.md)
    - [Layotto社区晋升规则](zh/community/promote.md)
    - [Layotto社区成员](zh/community/people.md)
- 博客
    - [蚂蚁云原生应用运行时的探索和实践 - ArchSummit 上海](zh/blog/exploration-and-practice-of-antcloud-native-application-runtime-archsummit-shanghai.md)
    - [MOSN子项目Layotto：开启服务网格+应用运行时新篇章](zh/blog/mosn-subproject-layotto-opening-a-new-chapter-in-service-grid-application-runtime/index.md)
    - 源码分析
        - [启动流程](zh/blog/code/start_process/start_process.md)
        - [处理 RPC 请求](zh/blog/code/layotto-rpc/index.md)
        - [WebAssembly 相关](zh/blog/code/webassembly/index.md)
        - [7层流量治理，接口限流](zh/blog/code/flowcontrol/flowcontrol_code_analyze.md)
        - [源码解析 4层流量治理，tcp流量dump](zh/blog/tcpcopy_code_analyze.md)
- [视频合集](zh/video/README.md)
