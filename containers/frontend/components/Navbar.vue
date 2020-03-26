<template>
  <nav class="navbar">
    <nuxt-link to="/" class="navbar-brand">
      Goonstone
    </nuxt-link>
    <div class="navbar-menu">
      <div class="navbar-start">
        <div class="navbar-item">
          <button class="button">
            <!-- TODO icon -->
            <i class="icon ion-md-add"></i>
            Submit a photo
          </button>
        </div>
        <span class="navbar-item">
          username
        </span>
        <div class="navbar-item">
          <nuxt-link to="/login" class="button button-link">
            Login / Register
          </nuxt-link>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
    export default {
        name: "Navbar"
    }
</script>

<style scoped>

</style>
