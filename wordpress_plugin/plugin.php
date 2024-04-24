<?php
/*
Plugin Name: WP Manager Plugin
Description: Provides endpoints to manage WordPress settings and data.
Version: 1.0
Author: Ashba22
License: GPL2
*/

defined( 'ABSPATH' ) || exit;
require_once( ABSPATH . 'wp-admin/includes/plugin.php' );

add_action( 'rest_api_init', function () {
    register_rest_route( 'wp-manager-plugin/v1', '/installed-plugins/', array(
        'methods'  => 'GET',
        'callback' => 'wp_manager_plugin_get_installed_plugins_data',
    ) );
} );
 

function wp_manager_plugin_get_installed_plugins_data( $request ) {
    // Validate API key
    if ( ! wp_manager_plugin_validate_api_key( $request ) ) {
        return new WP_Error( 'rest_forbidden', esc_html__( 'Invalid API Key', 'text-domain' ), array( 'status' => 401 ) );
    }


   
    $installed_plugins = get_plugins();


    $plugins_data = array();

    // Loop through each installed plugin
    foreach ( $installed_plugins as $plugin_path => $plugin_data ) {
        // Add plugin data to the array
        $plugins_data[] = array(
            'name'    => $plugin_data['Name'],
            'version' => $plugin_data['Version'],
            'author'  => $plugin_data['Author'],

        );
    }

    // Return the array of plugin data
    return rest_ensure_response( $plugins_data );
}


function wp_manager_plugin_validate_api_key( $request ) {
    $api_key = $request->get_param( 'api_key' );
    $stored_api_key = get_option( 'wp_manager_plugin_api_key' );
    return ! empty( $api_key ) && hash_equals( $stored_api_key, $api_key );
}

add_action( 'admin_menu', 'wp_manager_plugin_add_menu_item' );
function wp_manager_plugin_add_menu_item() {
    add_menu_page(
        'API Key Management',
        'API Key',
        'manage_options',
        'wp-manager-plugin-api-key',
        'wp_manager_plugin_api_key_page',
        'dashicons-lock',
        30
    );
}


function wp_manager_plugin_api_key_page() {
    
    if ( isset( $_POST['wp_manager_plugin_generate_api_key'] ) ) {
        // Generate a new API key
        $api_key = wp_generate_password( 64, false );

       
        update_option( 'wp_manager_plugin_api_key', $api_key );

    
        echo '<div class="updated"><p>New API key generated successfully: ' . esc_html( $api_key ) . '</p></div>';
    }

    $current_api_key = get_option( 'wp_manager_plugin_api_key' );
    ?>
    <div class="wrap">
        <h1>API Key Management</h1>
        <p>Current API key:</p>
        <p><strong><?php echo esc_html( $current_api_key ); ?></strong></p>
        <form method="post" action="">
            <input type="hidden" name="wp_manager_plugin_generate_api_key" value="1" />
            <input type="submit" class="button button-primary" value="Generate New Key" />
        </form>
    </div>
    <?php
}

function wp_manager_plugin_generate_api_key() {

    $api_key = wp_generate_password( 64, false );

    update_option( 'wp_manager_plugin_api_key', $api_key );
}


register_activation_hook( __FILE__, 'wp_manager_plugin_generate_api_key' );
