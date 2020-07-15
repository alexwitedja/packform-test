<template>
<div>
  <b-container class="bv-example-row">
    <b-row id="search-row" class="left-align">
      <b-col id="search" cols="2"><b-icon-search id="search-icon" /><strong>Search</strong></b-col>
      <b-col cols="10">
        <b-form-input
          v-model="filter"
          type="search"
          id="filterInput"
          placeholder="Search product/order name"
        ></b-form-input>
      </b-col>
    </b-row>
    
    <b-row cols="1" id="date-row" class="left-align">
      <b-col cols="12"><strong>Created Date</strong></b-col>
      <b-col cols="12"><date-range-picker v-model="range" @change="dateRange" :options="drpOptions" /></b-col>
    </b-row>

    <b-row cols="1" class="left-align">
      <b-col cols="12">Total Revenue: <strong>${{ total_revenue }}</strong></b-col>
    </b-row>
    
    <div style="margin-bottom:30px;">
      <b-table
        :items="items"
        :fields="fields"
        :per-page="perPage"
        :current-page="currentPage"
        :sort-by.sync="sortBy"
        :sort-desc.sync="sortDesc"
        sort-icon-right
        responsive="sm"
        :filter="filter"
        :filterIncludedFields="filterOn"
        id="my-table"
        @filtered="onFiltered"
        ref="table"
      >
        <template slot="order_name" slot-scope="{ value, item }">
          {{ item.Id }}
          <b-button @click="doSomthing(item.Id)">{{ item.Id }}</b-button>
        </template>
      </b-table>
    </div>
    
    <b-container class="bv-example-row">
      <b-row class="justify-content-md-center">
        <b-col cols="2">Pages: {{ pages }}</b-col>
        <b-col cols="2">
          <b-pagination
            v-model="currentPage"
            :total-rows="totalRows"
            :per-page="perPage"
            hide-goto-end-buttons
            aria-controls="my-table"
            align='center'
            limit=1
            id="pagination-button"
          ></b-pagination>
        </b-col>
        <b-col cols="2">Go to <b-form-input v-model="currentPage" :value="currentPage" id="go-to-input"></b-form-input></b-col>
      </b-row>
    </b-container>
  </b-container>
</div>
</template>

<script>
import axios from 'axios'
import moment from 'moment'
import { BIconSearch } from 'bootstrap-vue'

// Vue stuff here
export default {
  name: 'App',

  components:{
    BIconSearch,
  },

  data: function() {
    return {
      // Begin data for table.
      sortBy: '',
      sortDesc: false,
      fields: [
        { key: 'order_name', sortable: false,  }, 
        { key: 'company_name', sortable: false }, 
        { key: 'customer_name', sortable: false }, 
        { key: 'delivered_amount', 
          sortable: false, 
          formatter: value => {
            return (value > 0 ) ? '$' + value : '-'
          } 
        }, 
        { key: 'order_date', 
          sortable: true,
          formatter: value => {
            return this.dateFormatter(value)
          }
        }, 
        { key: 'total_amount',
          sortable: false,
          formatter: value => {
            return '$' + value
          }
        }
      ],
      items: [
        
      ],
      perPage: 5,
      currentPage: 1,
      filter: null,
      filterOn: ["order_name"],
      totalRows: 0,
      // End data for table.
      total_revenue: 0,
      // Date range picker data.
      range: [],
      drpOptions: {
        startDate: moment().startOf('year'),
        endDate: moment().endOf('year')
      }
    }
  },
  // For pagination
  computed: {
    customFilter() {
      let items = this.items 

      return items;
    }
  },
  // Created is called before DOM element is created.
  created() {
    axios.get("http://localhost:9999/api/orders", { crossdomain: true }).then(result => {
      this.items = result.data
      this.totalRows = this.items.length
      this.pages = Math.ceil(this.totalRows / this.perPage)
      this.totalRevenue(this.items)
    }).catch(error => {
      console.log(error)
    })
  },

  methods: {
    onFiltered(filteredItems) {
      // Trigger pagination to update the number of buttons/pages due to filtering
      this.totalRows = filteredItems.length
      this.currentPage = 1
      this.pages = Math.ceil(this.totalRows / this.perPage)
      this.totalRevenue(filteredItems)
    }, 
    totalRevenue(list) {
      var sum = 0;
      list.map(order => {
        sum += order.total_amount
      })

      this.total_revenue = sum.toFixed(2)
    },
    dateFormatter(date) {
      const d = new Date( Date.parse(date) )
      const mo = new Intl.DateTimeFormat('en', { month: 'short', timeZone: 'Australia/Sydney' }).format(d)
      const da = new Intl.DateTimeFormat('en', { day: 'numeric', timeZone: 'Australia/Sydney' }).format(d)
      const ti = new Intl.DateTimeFormat('en', { hour12: true, hour: 'numeric', minute: 'numeric',timeZone: 'Australia/Sydney' }).format(d)
      var daySuffix = 'st'
      if(da.charAt(da.length - 1) == '1') {
        daySuffix = 'st'
      } else if(da.charAt(da.length - 1) == '2') {
        daySuffix = 'nd'
      } else if(da.charAt(da.length - 1) == '3') {
        daySuffix = 'rd';
      } else {
        daySuffix = 'th'
      }

      const formatted = `${mo} ${da}${daySuffix}, ${ti}`
      return formatted
    },
    dateRange() {
      // Sends post request to backend and update this.items
      if (this.range.length > 0) {
        const splitStart = this.range[0].split("/")
        const splitEnd = this.range[1].split("/")
        const startDate = new Date(splitStart[2], splitStart[1], splitStart[0]).toISOString()
        const endDate = new Date(splitEnd[2], splitEnd[1], splitEnd[0]).toISOString()

        const payload = {
          start: startDate,
          end: endDate
        }

        axios({
          method: 'post',
          headers: { 'Accept': 'application/json',
                    'Content-Type': 'application/json;charset=UTF-8' },
          url: 'http://localhost:9999/api/ordersBetween',
          data: payload
        }).then(result => {
          this.items = result.data
          this.totalRows = this.items.length
          this.pages = Math.ceil(this.totalRows / this.perPage)
          this.totalRevenue(this.items)
          // Finally refreshes b-table
          this.$refs.table.refresh()
        });
      }
    }
  }
}
</script>

<style scoped>
  #search-icon {
    margin-right: 20px;  
  }

  #search {
    font-size: 25px
  }

  .col, .col-10 {
    padding: 0px;
  }

  .left-align {
    text-align: left;
    margin-bottom: 30px;
  }

  #date-row {
    margin-top: 10px;
    margin-bottom: 10px;
    text-align: left;
  }

  #go-to-input {
    display: inline;
    width: 25%;
  }

</style>