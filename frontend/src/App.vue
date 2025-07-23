<!--
 * @Author: JimZhang
 * @Date: 2025-07-22 17:29:27
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-23 11:08:10
 * @FilePath: /server/frontend/src/App.vue
 * @Description: 
 * 
-->
<script setup lang="ts">
Tools.localStorage.setItem('test', 'testasda')
</script>

<template>
  <div>
    {{ app.getConfig('baseUrl') }}
    {{ Tools.localStorage.getItem('test') }}
    <!-- {{ lpk }} -->
  </div>

</template>
