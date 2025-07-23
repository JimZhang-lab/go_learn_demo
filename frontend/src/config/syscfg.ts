export interface ISysCfgBModItem {
    name: string; // 模块名称
    enable: boolean; // 是否启用
}


export interface ISysCfg{

    baseUrl: string; // 后端主机地址和监听端口
    bmodNames: { // 业务模块列表
        name: string; // 模块名称
        enable: boolean; // 是否启用
    }[]
}

const iSysCfg: ISysCfg = {
    baseUrl: 'http://localhost:8080',
    bmodNames: [
        {
            name: 'mod_demo',
            enable: true
        }
    ]
}

export default iSysCfg;