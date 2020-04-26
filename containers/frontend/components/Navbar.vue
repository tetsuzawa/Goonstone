<template>
  <nav class="navbar">
    <nuxt-link to="/" class="navbar__brand">
      Goonstone
    </nuxt-link>
    <div class="navbar__menu">
      <div class="navbar-start">
        <div v-if="hasLoggedIn" class="navbar__item">
          <button class="button" @click="showForm = !showForm">
            <!-- TODO icon -->
            <i class="icon icon-md-add"></i>
            Submit a photo
          </button>
        </div>
        <span v-if="hasLoggedIn" class="navbar__item">
          {{ username }}
        </span>
        <div v-else class="navbar__item">
          <nuxt-link to="/login" class="button button--link">
            Login / Register
          </nuxt-link>
        </div>
      </div>
    </div>
<!--    <PhotoForm :value="showForm" @input="$emit('input', $event.target.showForm)"/>-->
    <PhotoForm v-model="showForm"/>
  </nav>
</template>

<script>
  import PhotoForm from "./PhotoForm";

    export default {
        name: "Navbar",
        components: {
            PhotoForm
        },
        data () {
          return{
              showForm: false
          }
        },
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
