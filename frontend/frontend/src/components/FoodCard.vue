<template>
  <div class="tile is-parent">
    <div class="tile is-child">
      <div class="nicebox">
        {{data}}
      </div>

    </div>
  </div>
</template>

<script>


export default {
  name: 'FoodCard',
  props: ['data'],
  data () {
    return {
    }
  },
  mounted: function(){

  }
}
</script>

<style scoped>
.nicebox{
  border-radius: 4px;
  padding: 20px;
  border:1px solid #E0E0EA;
  background-color: rgba(14, 105, 161, 0.03);
}
</style>
