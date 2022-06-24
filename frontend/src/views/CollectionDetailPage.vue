<template>
    <div class="col-md-10 offset-md-1 py-2">
        <div class="row">
            <div class="col-8">
                <h3 v-if="this.loading">Loading...</h3>
                <h3 v-else>
                    Collection: {{ this.collection.Name }}
                </h3>
            </div>
            <div class="col-4 text-right">
                <router-link :to="{ name: 'collection_list' }">
                    <button class="btn btn-secondary btn-sm">Back</button>
                </router-link>
                &nbsp;
                <router-link v-if="!this.loading" :to="{ name: 'collection_edit', params: { id: this.collection.ID } }">
                    <button class="btn btn-primary btn-sm">Update</button>
                </router-link>
            </div>
        </div>
        <div v-if="!this.loading" class="row">
            <div class="col-sm-12">
                <table class="table table-bordered">
                    <tr>
                        <th>Name</th>
                        <td>{{ this.collection.Name }}</td>
                    </tr>
                    <tr>
                        <th>Path Prefix</th>
                        <td>{{ this.collection.PathPrefix }}</td>
                    </tr>
                    <tr>
                        <th>Content Type</th>
                        <td>{{ this.collection.ContentType }}</td>
                    </tr>
                </table>
            </div>
        </div>
        <div class="row">
            <hr />
        </div>
        <div v-if="!this.loading" class="row">
            <div class="col-sm-12">
                <EndpointsTable :endpoints="this.collection.Endpoints" :collectionID="this.collection.ID" />
            </div>
        </div>
    </div>
</template>

<script>
import axios from "axios";
import EndpointsTable from "../components/EndpointsTable.vue";

export default {
    name: "CollectionDetailPage",

    components: {
        EndpointsTable,
    },

    data() {
        return {
            loading: false,
            collection: null,
            errorMessage: "",
        }
    },

    created() {
        this.$watch(
            () => this.$route.params, () => {
                this.fetchData();
            }, { immediate: true }
        );
    },

    methods: {
        fetchData() {
            this.errorMessage = null;
            this.collection = null;
            this.loading = true;

            if (this.$route.params.id != undefined) {
                axios.get("http://localhost:3000/itinerary/collection/"+this.$route.params.id)
                .then((response) => {
                    this.collection = response.data;
                    this.loading = false;
                })
                .catch((error) => {
                    this.errorMessage = error;
                })
            }
        }
    }
}
</script>