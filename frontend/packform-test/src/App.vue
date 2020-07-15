<template>
  <div id="app">
    <h1>Packform full stack test.</h1>
    <div id="nav">
      <router-link to="/">Home</router-link> | 
      <router-link to="/orders">Orders</router-link> | 
    </div>
    <router-view />
  </div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
