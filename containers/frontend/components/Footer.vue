<template>
  <footer class="footer">
    <button class="button button-link">Logout</button>
    <nuxt-link to="/login" class="button button-link">Login / Register</nuxt-link>
  </footer>
</template>

<script>
    export default {
        name: "Footer",
        methods: {
            async logout() {
                await this.$store.dispatch('auth/logout')
                this.$router.push('/login')
            }
        }
    }
</script>

<style scoped>

</style>
