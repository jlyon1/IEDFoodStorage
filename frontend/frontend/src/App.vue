<template>
  <Main></Main>
</template>

<script>
import 'bulma/css/bulma.css'
import Main from '@/components/Main'
export default {
  name: 'App',
  components: {
    Main
  }
}
</script>

<style lang="sass-loader">
@import url('https://fonts.googleapis.com/css?family=Open+Sans|Raleway');
$section-padding: 1rem;
$family-source-sans: "Raleway", sans-serif;
$family-primary: $family-source-sans;
@import '~bulma/bulma';
</style>
