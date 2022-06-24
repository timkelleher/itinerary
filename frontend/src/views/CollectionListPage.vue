<template>
    <div class="row">
        <div class="col-md-12 py-2">
            <div v-if="this.errorMessage">
                <div class="alert alert-danger" role="alert">
                    {{ this.errorMessage }}
                </div>
            </div>
            <div v-else>
                <CollectionsTable :collections="collections" />
            </div>
        </div>
    </div>
</template>

<script>
import axios from "axios";
import CollectionsTable from "../components/CollectionsTable.vue";

export default {
    name: "CollectionListPage.vue",

    components: {
        CollectionsTable,
    },

    data() {
        return {
            errorMessage: "",
            statuses: [],
            collections: [],
        }
    },

    methods: {
        getSets() {
            this.errorMessage = "";
            axios.get("http://localhost:3000/itinerary/collections")
                .then((response) => {
                    this.collections = response.data;
                })
                .catch((error) => {
                    this.errorMessage = error;
                })
        }  
    },

    mounted() {
        this.getSets();
    }
}
</script>
