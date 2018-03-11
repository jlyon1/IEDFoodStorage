<template>
  <div>
    <div class="header">
      <div class="">Team Maet Food Inventory</div>
    </div>
    <div class="section bdy">
      <div class="container">
        <div class="subtitle is-2">
          Your Pantry:
        </div>
        <div style="margin-bottom: 15px;">
          <button class="button">Add</button>
        </div>
        <div class="tile is-ancestor" v-for="i in Math.ceil(foodItems.Foods.length/3)">
          <FoodCard v-for="(item,j) in foodItems.Foods.slice((i - 1) * 3, i * 3)" v-bind:foodData="item" v-on:reload="reload"></FoodCard>
        </div>
      </div>
    </div>
    <div class="footer mt">
      <div class="container has-text-centered">
        Made with 🍖 by Taem Maet for IED Spring 2018
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
  .mt{
    margin-top: 50px;
  }
  .bdy{
    background-color:#fefefe;

  }
  .header{
    background-color: rgba(14, 105, 161, 0.03);
    padding: 20px;
    font-family: 'Open Sans', sans-serif;
    font-weight: 500;
    box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);

  }
</style>
