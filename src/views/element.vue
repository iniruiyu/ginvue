<template>

<input111/>

  <h4>1. 按钮</h4>
  <el-row class="mb-4">
    <el-button>Default</el-button>
    <el-button type="primary">Primary</el-button>
    <el-button type="success">Success</el-button>
    <el-button type="info">Info</el-button>
    <el-button type="warning">Warning</el-button>
    <el-button type="danger">Danger</el-button>
  </el-row>
  <hr />
  <h4>Plain 朴素的按钮</h4>
  <el-row class="mb-4">
    <el-button plain>Plain</el-button>
    <el-button type="primary" plain>Primary</el-button>
    <el-button type="success" plain>Success</el-button>
    <el-button type="info" plain>Info</el-button>
    <el-button type="warning" plain>Warning</el-button>
    <el-button type="danger" plain>Danger</el-button>
  </el-row>
  <hr />
  <h4>round 圆角按钮</h4>
  <el-row class="mb-4">
    <el-button round>Round</el-button>
    <el-button type="primary" round>Primary</el-button>
    <el-button type="success" round>Success</el-button>
    <el-button type="info" round>Info</el-button>
    <el-button type="warning" round>Warning</el-button>
    <el-button type="danger" round>Danger</el-button>
  </el-row>

  <hr />
  <h4>circle 圆形按钮</h4>
  <el-row>
    <el-button :icon="Search" circle />
    <el-button type="primary" :icon="Edit" circle />
    <el-button type="success" :icon="Check" circle />
    <el-button type="info" :icon="Message" circle />
    <el-button type="warning" :icon="Star" circle />
    <el-button type="danger" :icon="Delete" circle />
  </el-row>

  <h2>2.图标</h2>
  <h4>//您需要从 @element-plus/icons-vue 中导入所有图标并进行全局注册。</h4>
  <el-icon><Plus /></el-icon>
  <el-icon><Minus /></el-icon>
  <h4>2.1 大小和颜色</h4>
  <el-icon :size="30" :color="'red'"><Search /></el-icon>

  <h4>2.2 按钮中使用图标</h4>
  <el-button type="primary">
    <el-icon style="vertical-align: middle">
      <Search />
    </el-icon>
    <span style="vertical-align: middle"> | Search </span>
  </el-button>
  <el-button type="success">
    <el-icon style="vertical-align: middle">
      <Search />
    </el-icon>
  </el-button>

  <hr />
  <h4>2.3 按钮组合</h4>
  <el-button-group>
    <el-button type="primary" icon="el-icon-arrow-left">上一页</el-button>

    <el-button type="success">
      <el-icon style="vertical-align: middle">
        <Search />
      </el-icon>
    </el-button>

    <el-button type="success">
      <el-icon style="vertical-align: middle">
        <Calendar />
      </el-icon>
    </el-button>

    <el-button type="primary">下一页<el-icon icon="el-icon-arrow-right" /></el-button>
  </el-button-group>
</template>

<script lang="ts" setup>
import { Calendar, Check, Delete, Edit, Message, Search, Star } from '@element-plus/icons-vue'
import { FIRST_KEYS } from 'element-plus'

import input111 from '@/components/input.vue'

</script>
