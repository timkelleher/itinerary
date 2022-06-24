<template>
    <div class="row">
        <div class="col-md-10 offset-md-1 py-2">
            <div class="row">
                <div class="col-8 text-left">
                    <h3>New Collection</h3>
                </div>
                <div class="col-4 text-right">
                    <router-link :to="{ name: 'collection_list' }">
                        <button class="btn btn-secondary btn-sm">Back</button>
                    </router-link>
                    &nbsp;
                    <button @click="persistCollection" role="button" class="btn btn-primary btn-sm">Save</button>
                </div>
            </div>
            <div class="row" v-if="this.errorMessage">
                <div class="col-sm-12">
                    <div class="alert alert-danger" role="alert">
                        {{ this.errorMessage }}
                    </div>
                </div>
            </div>
            <div class="row" v-if="this.successMessage">
                <div class="col-sm-12">
                    <div class="alert alert-success" role="alert">
                        {{ this.successMessage }}
                    </div>
                </div>
            </div>
            <div>
                <form @submit="persistCollection">
                    <div class="row">
                        <div class="col-md-6">
                            <div class="form-group">
                                <label for="name">Name</label>
                                <input type="name" class="form-control" id="name" placeholder="Collection name" v-model="collection.Name" />
                            </div>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-md-6">
                            <div class="form-group">
                                <label for="name">Path Prefix</label>
                                <input type="name" class="form-control" id="path_prefix" placeholder="Path Prefix" v-model="collection.PathPrefix" />
                            </div>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-md-6">
                            <div class="form-group">
                                <label for="name">Content Type</label>
                                <input type="name" class="form-control" id="content_type" placeholder="Content Type" v-model="collection.ContentType" />
                            </div>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
import axios from "axios";

export default {
    name: "CollectionAddPage",

    data() {
        return {
            methods: [],
            responseStatusCodes: [],
            errorMessage: "",
            successMessage: "",
            cachedPathPrefix: "/",
            collection: {
                Name: "",
                PathPrefix: "",
                ContentType: ""
            }
        }
    },

    mounted() {
        
    },

    computed: {
        
    },

    methods: {
        persistCollection() {
            axios.put("http://localhost:3000/itinerary/collection", {
                "name": this.collection.Name,
                "path_prefix": this.collection.PathPrefix,
                "content_type": this.collection.ContentType,
            })
                .then((response) => {
                    this.collection = response.data;
                    this.successMessage = "Collection has been created";
                    this.$router.push({
                        name: "collection_detail",
                        params: {"id": this.collection.ID},
                    })
                })
                .catch((error) => {
                    this.errorMessage = error;
                })
        }
    }
}
</script>