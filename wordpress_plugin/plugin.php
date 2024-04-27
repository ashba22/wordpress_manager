<?php
/*
Plugin Name: WP Manager Plugin
Description: Provides endpoints to manage WordPress settings and data.
Version: 1.1
Author: Ashba22
License: GPL2
*/

defined( 'ABSPATH' ) || exit;
require_once( ABSPATH . 'wp-admin/includes/plugin.php' );
require_once( ABSPATH . 'wp-admin/includes/theme.php' );
require_once( ABSPATH . 'wp-admin/includes/update.php' );
require_once( ABSPATH . 'wp-admin/includes/post.php' );
require_once( ABSPATH . 'wp-admin/includes/user.php' );
require_once( ABSPATH . 'wp-admin/includes/media.php' );

add_action( 'rest_api_init', function () {
    register_rest_route( 'wp-manager-plugin/v1', '/installed-plugins/', array(
        'methods'  => 'GET',
        'callback' => 'wp_manager_plugin_get_installed_plugins_data',
    ) );
    register_rest_route( 'wp-manager-plugin/v1', '/installed-themes/', array(
        'methods'  => 'GET',
        'callback' => 'wp_manager_plugin_ger_installed_themes_data',
    ) );
    register_rest_route( 'wp-manager-plugin/v1', '/website-info/', array(
        'methods'  => 'GET',
        'callback' => 'wp_manager_all_data',
    ) );
    /// add posts,pages,users,media,comments
    register_rest_route( 'wp-manager-plugin/v1', '/posts/', array(
        'methods'  => 'GET',
        'callback' => 'wp_manager_all_posts',
    ) );
    register_rest_route( 'wp-manager-plugin/v1', '/pages/', array(
        'methods'  => 'GET',
        'callback' => 'wp_manager_all_pages',
    ) );
    register_rest_route( 'wp-manager-plugin/v1', '/users/', array(
        'methods'  => 'GET',
        'callback' => 'wp_manager_all_users',
    ) );
    register_rest_route( 'wp-manager-plugin/v1', '/media/', array(
        'methods'  => 'GET',
        'callback' => 'wp_manager_all_media',
    ) );
    register_rest_route( 'wp-manager-plugin/v1', '/comments/', array(
        'methods'  => 'GET',
        'callback' => 'wp_manager_all_comments',
    ) );
    register_rest_route( 'wp-manager-plugin/v1', '/updates/', array(
        'methods'  => 'GET',
        'callback' => 'wp_manager_all_updates',
    ) );
    
} );


function wp_manager_all_updates($request) {
    // Validate API key
    if ( ! wp_manager_plugin_validate_api_key( $request ) ) {
        return new WP_Error( 'rest_forbidden', esc_html__( 'Invalid API Key', 'text-domain' ), array( 'status' => 401 ) );
    }

    $updates = get_core_updates();
    $plugin_updates = get_plugin_updates();
    $theme_updates = get_theme_updates();




    $data = array(
        'core'   => $updates,
        'plugins' => $plugin_updates,
        'themes'  => $theme_updates,
    );

    return rest_ensure_response( $data );
}



function wp_manager_all_comments($request) {
    // Validate API key
    if ( ! wp_manager_plugin_validate_api_key( $request ) ) {
        return new WP_Error( 'rest_forbidden', esc_html__( 'Invalid API Key', 'text-domain' ), array( 'status' => 401 ) );
    }

    $comments = get_comments();

    $comments_data = array();

    // Loop through each comment
    foreach ( $comments as $comment ) {
        // Add comment data to the array
        $comments_data[] = array(
            'comment_author'    => $comment->comment_author,
            'comment_content'   => $comment->comment_content,
            'comment_date'      => $comment->comment_date,
            'comment_post_title' => get_the_title($comment->comment_post_ID),
        );
    }

    // Return the array of comment data
    return rest_ensure_response( $comments_data );
}

function wp_manager_all_media($request) {
    // Validate API key
    if ( ! wp_manager_plugin_validate_api_key( $request ) ) {
        return new WP_Error( 'rest_forbidden', esc_html__( 'Invalid API Key', 'text-domain' ), array( 'status' => 401 ) );
    }

    $media = get_posts( array(
        'post_type' => 'attachment',
        'posts_per_page' => -1,
    ) );

    $media_data = array();

    // Loop through each media item
    foreach ( $media as $media_item ) {
        // Add media data to the array
        $media_data[] = array(
            'media_title'    => $media_item->post_title,
            'media_url'      => $media_item->guid,
            'media_date'     => $media_item->post_date,
        );
    }

    // Return the array of media data
    return rest_ensure_response( $media_data );
}

function wp_manager_all_users($request) {
    // Validate API key
    if ( ! wp_manager_plugin_validate_api_key( $request ) ) {
        return new WP_Error( 'rest_forbidden', esc_html__( 'Invalid API Key', 'text-domain' ), array( 'status' => 401 ) );
    }

    $users = get_users();

    $users_data = array();

    // Loop through each user
    foreach ( $users as $user ) {
        // Add user data to the array
        $users_data[] = array(
            'user_login'    => $user->user_login,
            'user_email'    => $user->user_email,
            'user_registered' => $user->user_registered,
        );
    }

    // Return the array of user data
    return rest_ensure_response( $users_data );
}

function wp_manager_all_pages($request) {
    // Validate API key
    if ( ! wp_manager_plugin_validate_api_key( $request ) ) {
        return new WP_Error( 'rest_forbidden', esc_html__( 'Invalid API Key', 'text-domain' ), array( 'status' => 401 ) );
    }

    $pages = get_posts( array(
        'post_type' => 'page',
        'posts_per_page' => -1,
    ) );

    $pages_data = array();

    // Loop through each page
    foreach ( $pages as $page ) {
        // Add page data to the array
        $pages_data[] = array(
            'page_title'    => $page->post_title,
            'page_content'  => $page->post_content,
            'page_date'     => $page->post_date,
        );
    }

    // Return the array of page data
    return rest_ensure_response( $pages_data );
}

function wp_manager_all_posts($request) {
    // Validate API key
    if ( ! wp_manager_plugin_validate_api_key( $request ) ) {
        return new WP_Error( 'rest_forbidden', esc_html__( 'Invalid API Key', 'text-domain' ), array( 'status' => 401 ) );
    }

    $posts = get_posts( array(
        'post_type' => 'post',
        'posts_per_page' => -1,
    ) );

    $posts_data = array();

    // Loop through each post
    foreach ( $posts as $post ) {
        // Add post data to the array
        $posts_data[] = array(
            'post_title'    => $post->post_title,
            'post_content'  => $post->post_content,
            'post_date'     => $post->post_date,
        );
    }

    // Return the array of post data
    return rest_ensure_response( $posts_data );
}




function wp_manager_all_data($request) {
    // Validate API key
    if ( ! wp_manager_plugin_validate_api_key( $request ) ) {
        return new WP_Error( 'rest_forbidden', esc_html__( 'Invalid API Key', 'text-domain' ), array( 'status' => 401 ) );
    } 

    $wordpressVersion = get_bloginfo('version');
    $phpVersion = phpversion();
    $mysqlVersion = $GLOBALS['wpdb']->db_version();
    $serverInfo = $_SERVER['SERVER_SOFTWARE'];
    $websiteTitle = get_bloginfo('name');
    $websiteUrl = get_bloginfo('url');
    $websiteEmail = get_bloginfo('admin_email');
    $websiteLanguage = get_bloginfo('language');
    $websiteCharset = get_bloginfo('charset');

    $data = array(
        'Website Title' => $websiteTitle,
        'Website URL' => $websiteUrl,
        'Website Email' => $websiteEmail,
        'Website Language' => $websiteLanguage,
        'Website Charset' => $websiteCharset,
        'WordPress Version' => $wordpressVersion,
        'PHP Version' => $phpVersion,
        'MySQL Version' => $mysqlVersion,
        'Server Info' => $serverInfo,
    );
 

    return rest_ensure_response( $data );
}

 
function wp_manager_plugin_ger_installed_themes_data( $request ) {
    // Validate API key
    if ( ! wp_manager_plugin_validate_api_key( $request ) ) {
        return new WP_Error( 'rest_forbidden', esc_html__( 'Invalid API Key', 'text-domain' ), array( 'status' => 401 ) );
    }

    $installed_themes = wp_get_themes();

    $themes_data = array();
    

    // Loop through each installed theme
    foreach ( $installed_themes as $theme ) {
        // Add theme data to the array
        $themes_data[] = array(
            'name'    => $theme->get( 'Name' ),
            'version' => $theme->get( 'Version' ),
            'author'  => $theme->get( 'Author' ),
        );
    }

    // Return the array of theme data
    return rest_ensure_response( $themes_data );
}

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
