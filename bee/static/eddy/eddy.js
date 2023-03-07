/**
 * Mots clés du dictionnaire_EDDY
 */

var $eddy_dico = [
  // Actions
  { text: 'checkbox:\n  get-sql: "select..."\n  set-sql: "update..."\n  alias-db: "mydb"', displayText: 'a_checkbox' },
  { text: 'group: ', displayText: 'a_group' },
  { text: 'hide-sqlite: true', displayText: 'a_hide-sqlite' },
  { text: 'hide: true', displayText: 'a_hide' },
  { text: 'label: ""', displayText: 'a_label' },
  { text: 'plugin: ""', displayText: 'a_plugin' },
  { text: 'sql: ""', displayText: 'a_sql' },
  { text: 'url: ""', displayText: 'a_url' },
  { text: 'with-confirm: true', displayText: 'a_with-confirm' },
  // elements:
  { text: 'args:\n    arg: "value"', displayText: 'e_args' },
  { text: 'class-sqlite: "select..."', displayText: 'e_class-sqlite' },
  { text: 'col-align: left center right', displayText: 'e_col-align' },
  { text: 'col-nowrap: true', displayText: 'e_col-nowrap' },
  { text: 'compute-sqlite: "select..."', displayText: 'e_compute-sqlite' },
  { text: 'dataset:\n        arg: "value"', displayText: 'e_dataset' },
  { text: 'default-sqlite: "select..."', displayText: 'e_default-sqlite' },
  { text: 'default: "value"', displayText: 'e_default' },
  { text: 'format-sqlite: "select strftime(\'%H:%M:%S\', {Milliseconds}/1000, \'unixepoch\')"', displayText: 'e_format-sqlite' },
  { text: 'group: admin', displayText: 'e_group' },
  { text: 'help: "Help me"', displayText: 'e_help' },
  { text: 'hide-on-mobile: true', displayText: 'e_hide-on-mobile' },
  { text: 'hide-sqlite: "select \'notnull\'"', displayText: 'e_hide-sqlite' },
  { text: 'hide: true', displayText: 'e_hide' },
  { text: 'items-sql: "SELECT ptf_id as \'key\', ptf_name as \'label\' From ptf order by ptf_name"', displayText: 'e_items-sql' },
  { text: 'items:\n    - key: ""\n    label: ""', displayText: 'e_items' },
  { text: 'jointure:\n    params:\n    join: "left outer join ptf on ptf_id = orders_ptf_id"\n      column: "ptf.ptf_quote"', displayText: 'e_jointure' },
  { text: 'label-long: "Numéro"', displayText: 'e_label-long' },
  { text: 'label-short: "N°"', displayText: 'e_label-short' },
  { text: 'max-length: 10', displayText: 'e_max-length' },
  { text: 'max: 50 # valeur max', displayText: 'e_max' },
  { text: 'min-length: 10', displayText: 'e_min-length' },
  { text: 'min: 3  # valeur min', displayText: 'e_min' },
  { text: 'order: 10', displayText: 'e_order' },
  { text: 'params:\n    p_', displayText: 'e_params' },
  { text: 'pattern: "[a-zA-Z0-9_-]+" "[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$"', displayText: 'e_pattern' },
  { text: 'place-holder: "texte quand vide"', displayText: 'e_place-holder' },
  { text: 'post-action: \n    - a_', displayText: 'e_post-action' },
  { text: 'protected: true', displayText: 'e_protected' },
  { text: 'read-only: true', displayText: 'e_read-only' },
  { text: 'required: true', displayText: 'e_required' },
  { text: 'type: t_', displayText: 'e_type' },
  { text: 'width: s m l xl max', displayText: 'e_width' },
  // forms:
  { text: 'actions:\n    - a_', displayText: 'f_actions' },
  { text: 'check-sqlite:\n    - "select if error..."', displayText: 'f_check-sqlite' },
  { text: 'join: ""', displayText: 'p_join (jointure)' },
  { text: 'group: ""', displayText: 'f_group' },
  { text: 'hide-submit: true', displayText: 'f_hide-submit' },
  { text: 'icon-name: ', displayText: 'f_icon-name' },
  { text: 'post-sql:\n    - "update..."', displayText: 'f_post-sql' },
  { text: 'title: ""', displayText: 'f_title' },
  // Params
  { text: 'actions:\n    - a_', displayText: 'p_actions' },
  { text: 'dataset: ""', displayText: 'p_dataset' },
  { text: 'column: ""', displayText: 'p_column (jointure)' },
  { text: 'description: ""', displayText: 'p_description' },
  { text: 'extra: ""', displayText: 'p_extra' },
  { text: 'form: ""', displayText: 'p_form (card)' },
  { text: 'header: ""', displayText: 'p_header' },
  { text: 'join: ""', displayText: 'p_join (jointure)' },
  { text: 'meta: ""', displayText: 'p_meta' },
  { text: 'path: ""', displayText: 'p_path' },
  { text: 'sql: ""', displayText: 'p_sql' },
  { text: 'src: ""', displayText: 'p_src' },
  { text: 'url: ""', displayText: 'p_url' },
  { text: 'table: ""', displayText: 'p_table' },
  { text: 'title: ""', displayText: 'p_title' },
  { text: 'target: ""', displayText: 'p_target (url)' },
  { text: 'view: ""', displayText: 'p_view (card)' },
  { text: 'vhere: ""', displayText: 'p_where' },
  { text: 'width: s m l xl max', displayText: 'p_width' },
  { text: 'with-confirm: true', displayText: 'p_with-confirm' },
  { text: 'with-image-editor: true', displayText: 'p_with-image-editor' },
  { text: 'with-input-file: true', displayText: 'p_with-input-file' },
  { text: 'with-input: true', displayText: 'p_with-input' },
  // Types
  { text: 'action', displayText: 't_action' },
  { text: 'amount', displayText: 't_amount' },
  { text: 'card', displayText: 't_card' },
  { text: 'button', displayText: 't_button' },
  { text: 'checkbox', displayText: 't_ckeckbox' },
  { text: 'counter: ""', displayText: 't_counter' },
  { text: 'date', displayText: 't_date' },
  { text: 'epdf', displayText: 't_pdf' },
  { text: 'email', displayText: 't_email' },
  { text: 'float', displayText: 't_float' },
  { text: 'image', displayText: 't_image' },
  { text: 'jointure', displayText: 't_jointure' },
  { text: 'list', displayText: 't_list' },
  { text: 'number', displayText: 't_number' },
  { text: 'password', displayText: 't_password' },
  { text: 'pdf', displayText: 't_pdf' },
  { text: 'percent', displayText: 't_percent' },
  { text: 'plugin', displayText: 't_plugin' },
  { text: 'radio', displayText: 't_radio' },
  { text: 'tag', displayText: 't_tag' },
  { text: 'tel', displayText: 't_tel' },
  { text: 'text', displayText: 't_text' },
  { text: 'textarea', displayText: 't_textarea' },
  { text: 'time', displayText: 't_time' },
  { text: 'url', displayText: 't_url' },
  // views:
  { text: 'actions:\n    - a_', displayText: 'v_actions' },
  { text: 'card:\n      header:\n        - r_\n      meta:\n       - r_\n      description:\n        - r_\n      extra:\n        - r_\n      footer:\n        - r_', displayText: 'v_mask' },
  { text: 'class-sqlite: ', displayText: 'v_class-sqlite' },
  { text: 'deletable: true', displayText: 'v_deletable' },
  { text: 'elements:\n  r_', displayText: 'v_elements' },
  { text: 'footer-sql: ""', displayText: 'v_footer-sql' },
  { text: 'form-add: ', displayText: 'v_form-add' },
  { text: 'form-edit: ', displayText: 'v_form-edit' },
  { text: 'form-view: ', displayText: 'v_form-view' },
  { text: 'group: ', displayText: 'v_group' },
  { text: 'hide-on-mobile: true', displayText: 'v_hide-on-mobile' },
  { text: 'hide: true', displayText: 'v_hide' },
  { text: 'icon-name: ', displayText: 'v_icon-name' },
  { text: 'limit: 50', displayText: 'v_limit' },
  { text: 'order-by: ""', displayText: 'v_order-by' },
  { text: 'post-sql:\n    - "update delete..."', displayText: 'v_post-sql' },
  { text: 'pre-update-sql: \n    - "update..."', displayText: 'v_pre-update-sql' },
  { text: 'title: ""', displayText: 'v_title' },
  { text: 'type: "card image table"', displayText: 'v_type' },
  { text: 'search: ""', displayText: 'v_search' },
  { text: 'with-line-number: true', displayText: 'v_with-line-number' },
];
