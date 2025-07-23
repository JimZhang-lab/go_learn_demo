/*
 * @Author: JimZhang
 * @Date: 2025-07-22 20:10:55
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-22 23:29:54
 * @FilePath: /server/frontend/src/config/init.ts
 * @Description: 
 * 
 */
import app from './app'
import Tools from '@/utils/Tools'

// 声明全局相关的类型
type IGlobalVarKey = 'app' | 'Ipk' | 'Tools' | 'Ajax'
type IGlobalVar = {
    [key in IGlobalVarKey]?: any
}

const iGlobalVar: IGlobalVar = {
    app, // 全局应用对象，包含一些全局数据与操作的方法
    Ipk: {},
    Tools,
    Ajax: {},
}

Object.keys(iGlobalVar).forEach(key => {
    // @ts-ignore
    window[key as IGlobalVarKey] = iGlobalVar[key as IGlobalVarKey]
})

export const initApp = async () => {

}