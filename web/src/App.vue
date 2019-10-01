<template>
  <v-app>
    <v-app-bar app dark color="primary">
      <v-app-bar-nav-icon>
        <v-icon>fa-database</v-icon>
      </v-app-bar-nav-icon>
      <v-toolbar-title>User CRM</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-toolbar-items>
        <v-btn text :to="{name: 'users'}" :exact="true">
          <v-icon left>fa-user-circle</v-icon>
          <span>
            Users
          </span>
        </v-btn>

        <v-btn text :to="{name: 'countries'}">
          <v-icon left>fa-globe-americas</v-icon>
          <span>
            Countries
          </span>
        </v-btn>

        <v-btn text href="https://github.com/memclutter/user-crm" target="_blank">
          <v-icon left>fab fa-github</v-icon>
          <span>
            Source
          </span>
        </v-btn>
      </v-toolbar-items>
    </v-app-bar>
    <v-content>
      <router-view></router-view>
    </v-content>
  </v-app>
</template>