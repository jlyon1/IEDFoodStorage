<template>
  <div>
    <div class="header">
      <div class="">Team Maet Food Inventory</div>
    </div>
    <div class="section">
      <div class="container">
        <FoodCard v-for="item in foodItems.Foods" v-bind:foodData="item" v-on:reload="reload"></FoodCard>
      </div>
    </div>
  </div>
</template>

<script>
import FoodCard from '@/components/FoodCard'


export default {

  name: 'Main',
  components: {FoodCard},
  data () {
    return {
      foodItems: {}
    }
  },
  methods: {

    reload: function(){
      let el = this;
      fetch("http://127.0.0.1:8081/").then(function(data){
        return data.json();
      }).then(function (resp){
        el.foodItems = resp
      });
    }
  },
  mounted: function(){
    let el = this;
    fetch("http://127.0.0.1:8081/").then(function(data){
      return data.json();
    }).then(function (resp){
      el.foodItems = resp
    });
  }
}
</script>

<style scoped>
.header{
  background-color: rgba(14, 105, 161, 0.03);
  padding: 20px;
  font-family: 'Open Sans', sans-serif;
  font-weight: 500;
}
</style>
