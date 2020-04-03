<template>
  <nav class="navbar">
    <nuxt-link to="/" class="navbar-brand">
      Goonstone
    </nuxt-link>
    <div class="navbar-menu">
      <div class="navbar-start">
        <div v-if="hasLoggedIn" class="navbar-item">
          <button class="button">
            <!-- TODO icon -->
            <i class="icon ion-md-add"></i>
            Submit a photo
          </button>
        </div>
        <span v-if="hasLoggedIn" class="navbar-item">
          username
        </span>
        <div v-else class="navbar-item">
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
        name: "Navbar",
        computed: {
            hasLoggedIn() {
                return this.$store.getters['auth/hasLoggedIn']
            },
            username() {
                return this.$store.getters['auth/username']
            }
        }
    }
</script>

<style scoped>

</style>
