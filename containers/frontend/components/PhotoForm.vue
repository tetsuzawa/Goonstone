<template>
  <div v-show="value" class="photo-form">
    <h2 class="title">Submit a photo</h2>
    <form class="form" @submit.prevent="submit">>
      <div class="errors" v-if="errors">
        <ul v-if="errors.photo">
          <li v-for="msg in errors.photo" :key="msg">{{ msg }}</li>
        </ul>
      </div>
      <input class="form__item" type="file" @change="onFileChange" >
      <output class="form-output" v-if="preview">
        <img :src="preview" alt="">
      </output>
      <div class="form__button">
        <button type="submit" class="button button-inverse" :disabled="!photo">submit</button>
      </div>
    </form>
  </div>
</template>

<script>
    import {STATUS_CREATED, STATUS_UNPROCESSABLE_ENTITY} from "../static/js/errorCodes";

    export default {
        name: "PhotoForm",
        props: {
            value: {
                type: Boolean,
                required: true
            }
        },
        data() {
            return {
                preview: null,
                photo: null,
                errors: null
            }
        },
        methods: {
            onFileChange(event) {
                if (event.target.files.length === 0) {
                    this.reset();
                    return false
                }
                if (!event.target.files[0].type.match('image.*')) {
                    this.reset();
                    return false
                }
                const reader = new FileReader();
                reader.onload = e => {
                    this.preview = e.target.result
                };
                reader.readAsDataURL(event.target.files[0]);
                this.photo = event.target.files[0];
            },
            reset() {
                this.preview = '';
                this.photo = null;
                this.$el.querySelector('input[type="file"]').value = null
            },
            async submit() {
                const formData = new FormData();
                formData.append('photo', this.photo);
                const response = await this.$axios.post('/photos', formData);

                if (response.status === STATUS_UNPROCESSABLE_ENTITY) {
                    this.errors = response.data.errors;
                    return false
                }

                this.reset();
                this.$emit('input', false);

                if (response.status !== STATUS_CREATED) {
                    this.$store.commit('error/setCode', response.status);
                    return false
                }

                this.$router.push(`/photos/${response.data.photo.id}`)
            }
        }
    }
</script>

<style scoped>

</style>
