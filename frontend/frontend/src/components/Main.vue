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
          <button v-if="en" class="button" @click="add">Add</button>
          <button v-if="en" class="button" @click="disable">Disable</button>
          <button v-else class="button" @click="enable">Enable</button>
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
import editor from 'vue2-medium-editor'
import 'bulma-extensions/bulma-calendar/dist/bulma-calendar.min.js'

export default {

  name: 'Main',
  components: {
    FoodCard,
    'medium-editor': editor
  },
  data () {
    return {
      foodItems: [],
      en: true,
    }
  },
  methods: {
    add: function(){
      var tmp = {
        "ID": -1,
        "Name": "Unknown",
        "ExpirationDate": "2018-03-10T00:00:00Z",
        "Position": -1,
        "PadNum": -1,
        "Count": 1,
        "Created":Date.now(),
      }
      this.foodItems.Foods.unshift(tmp)
      console.log(this.foodItems)
    },
    enable: function(){
      var val = {"Enabled": true}
      let el =  this
      fetch('/enabled', {
        method: 'post',
        body: JSON.stringify(val)
      }).then(function(response) {
        return response.text();
      }).then(function(data) {
        el.en = true
      });
    },
    disable: function(){
      var val = {"Enabled": false}
      let el =this;
      fetch('/enabled', {
        method: 'post',
        body: JSON.stringify(val)
      }).then(function(response) {
        return response.text();
      }).then(function(data) {
        el.en = false
      });
    },
    reload: function () {
      let el = this
      //el.foodItems = {}
      fetch('/get').then(function (data) {
        return data.json()
      }).then(function (resp) {
        if (el.foodItems[0].ID != -1){
          el.foodItems = resp
        }
        
      })
    }
  },
  mounted: function () {
    let el = this
    setInterval(el.reload,10000)
    fetch('/get').then(function (data) {
      // console.log(data.json())
      return data.json()
    }).then(function (resp) {
      console.log(resp)
      el.foodItems = resp
    })
    fetch('/enabled').then(function (data) {
      // console.log(data.json())
      return data.json()
    }).then(function (resp) {
      console.log(resp)
      el.en = resp.Enabled
    })
  }
}
</script>

<style scoped>

  .medium-editor-action {
    display: none;
  }
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
