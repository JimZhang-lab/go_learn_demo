/*
 * @Author: JimZhang
 * @Date: 2025-07-22 23:42:19
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-23 00:01:24
 * @FilePath: /server/frontend/src/config/lpk.ts
 * @Description: 
 * 
 */

const tblLpk: Record<string, string | string[]> = {}
export const initLpk = () => { 
    mergeLpk(import.meta.glob('@/locales/*.ts', {eager: true}))
}

export const getLocale = () => {
    return 'zh-CN'

}
type ILpkFile = {
    [path: string] :{
        default: Record<string, string | string[]>
    }
}
// {
//     {
//         'zh-CN.ts'
//     }
// }
// type IFnMergeLpk = (importLpkFiles: ILpkFile) => void
export const mergeLpk = (importLpkFiles: ILpkFile) => {
    const stLocaleLanguage = getLocale()
    for (const path in importLpkFiles){
        if (-1 == path.indexOf(stLocaleLanguage)){
            continue
        }else{

        }

        const { default: iLpkFileItem } = importLpkFiles[path]

        for (const stLpkKey in iLpkFileItem){
            tblLpk[stLpkKey] = iLpkFileItem[stLpkKey]
        }

    }
}

export const lpk = () => {


}

export const changeLocale = () => {


}