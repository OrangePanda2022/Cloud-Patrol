<script setup lang="ts"></script>

<template>
  <HeaderView>
    <template #page-content>
      <main>
        <RouterView />
      </main>
    </template>
  </HeaderView>
  <FooterView />
</template>
<style>
@import 'tailwindcss';
@plugin "daisyui";
</style>
