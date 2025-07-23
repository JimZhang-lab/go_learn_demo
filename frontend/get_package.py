'''
Author: JimZhang
Date: 2025-07-22 18:36:47
LastEditors: 很拉风的James
LastEditTime: 2025-07-22 19:15:16
FilePath: /server/frontend/get_package.py
Description: 

'''


package_dit = {
    "dependencies": {
    "@heroicons/vue": "^2.2.0",
    "@tailwindcss/vite": "^4.1.11",
    "axios": "^1.10.0",
    "daisyui": "^5.0.46",
    "js-cookie": "^3.0.5",
    "lodash": "^4.17.21",
    "nanoid": "^5.1.5",
    "normalize.css": "^8.0.1",
    "pinia": "^3.0.3",
    "qs": "^6.14.0",
    "vue": "^3.5.17",
    "vue-router": "^4.5.1"
  },
  "devDependencies": {
    "@tailwindcss/postcss": "^4.1.11",
    "@types/js-cookie": "^3.0.6",
    "@types/lodash": "^4.17.20",
    "@types/node": "^24.0.15",
    "@types/qs": "^6.14.0",
    "@vitejs/plugin-vue": "^6.0.0",
    "@vue/tsconfig": "^0.7.0",
    "autoprefixer": "^10.4.21",
    "postcss": "^8.5.6",
    "postcss-modules": "^6.0.1",
    "postcss-pxtorem": "^6.1.0",
    "tailwindcss": "^4.1.11",
    "typescript": "~5.8.3",
    "vite": "^7.0.4",
    "vue-tsc": "^2.2.12"
  }
}

package_name = []
for k, v in package_dit.items():
    # print(k)
    for k1, v1 in v.items():
        # print(k1, v1)
        package_name.append(k1)
        # print()
        
# import json


print(' '.join(package_name))