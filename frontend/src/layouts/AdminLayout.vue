<template>
  <q-layout view="hHh lpR fFf">
    <q-toolbar class="bg-primary glossy text-white">
      <q-btn flat round dense icon="menu" class="q-mr-sm" />
      <q-toolbar-title>TMA Starter Kit</q-toolbar-title>
      <q-btn flat round dense icon="whatshot" />
    </q-toolbar>
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>
<script setup lang="ts">
</script>
