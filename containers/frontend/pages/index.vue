<template>
  <div>
    <header>
      <Navbar/>
    </header>
    <main>
      <div class="container"></div>
    </main>
    <Footer/>
  </div>
</template>

<script>
import Navbar from "../components/Navbar";
import Footer from "../components/Footer";
import {STATUS_INTERNAL_SERVER_ERROR} from "../static/js/errorCodes";

export default {
    components: {
        Navbar,
        Footer
    },
    middleware: 'user',
    computed: {
        errorCode() {
            if (this.$store.state.error.code === STATUS_INTERNAL_SERVER_ERROR) {
                this.$router.push('/error/err500')
            }
            return this.$store.state.error.code
        }
    },
}
</script>

<style></style>
