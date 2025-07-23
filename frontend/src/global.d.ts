/*
 * @Author: JimZhang
 * @Date: 2025-07-22 20:07:06
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-22 23:35:30
 * @FilePath: /server/frontend/src/global.d.ts
 * @Description: 
 * 
 */


/**
 * 全局类型定义模块
 * 定义了全局的GlobalType命名空间和Window接口扩展
 * 包含两个主要类型定义：Ikey和IRecord
 */
import {IApp} from "@/config/app";
import {ITools} from "@/utils/Tools";

declare global {
    /**
     * 全局类型命名空间
     * 包含通用类型定义
     */
    declare namespace GlobalType{
        /**
         * 键类型定义
         * 支持字符串或数字类型的键
         */
        type Ikey = string | number;
        
        /**
         * 通用记录类型
         * 使用Ikey作为键的Record类型
         */
        type IRecord = Record<Ikey, any>;
    }
    const app: IApp;
    const Tools: ITools;
    /**
     * Window接口扩展
     * 添加app属性到全局Window接口
     */
    interface Window{
        /**
         * 全局应用对象
         * 类型为GlobalType.IRecord
         */
        app: IApp; // 全局应用对象，挂载一些全局数据与操作
        Tools: ITools; // 全局公用方法，包含一些公用方法
    }
}

/**
 * Vue组件属性扩展
 * 为Vue组件的自定义属性添加app属性定义
 */
declare module 'vue' {
    /**
     * 组件自定义属性扩展
     * 添加app属性到Vue组件自定义属性
     */
    interface ComponentCustomProperties {
        /**
         * 组件上下文应用对象
         * 类型为GlobalType.IRecord
         */
        app: IApp;
        Tools: ITools;
    }
}
export {}