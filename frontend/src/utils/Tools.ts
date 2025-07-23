import Cookies from "js-cookie"
import { set } from "lodash"

/*
 * @Author: JimZhang
 * @Date: 2025-07-22 23:04:41
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-22 23:14:45
 * @FilePath: /server/frontend/src/utils/Tools.ts
 * @Description: 
 * 
 */


const iTools = {
    Router: { // 路由操作命名空间


    },
    Store: { // 状态管理操作命名空间


    },
    localStorage: { // 本地存储命名空间
        // 设置localStorage
        setItem(key: string, value: string){
            localStorage.setItem(key, JSON.stringify(value))
        },
        // 获取localStorage
        getItem(key: string){
            const stValue = localStorage.getItem(key)
            try{
                return JSON.parse(stValue as string)
            } catch(e){
                return stValue
            }
        },
        // 删除localStorage
        removeItem(key: string){
            localStorage.removeItem(key)
        },

    },
    Cookies: { // cookie操作命名空间
        setItem(key: string, value: any, expire?: number){
            if (expire) {
                Cookies.set(key, JSON.stringify(value), { expires: expire })
            }else{
                // 单位为天
                Cookies.set(key, JSON.stringify(value), {expires: 7})
            }
        },
        getItem(key: string, defaultValue?: any){
            const stValue = Cookies.get(key) || defaultValue
            try{
                return JSON.parse(stValue)
            } catch (error) {
                return stValue

            }
        },
        removeItem(key: string){
            Cookies.remove(key)
        },

    },
    Time:{ // 日期时间操作命名空间

    },
    Dom:{ // DOM 元素操作命名空间


    }
}
export type ITools = typeof iTools
export default iTools