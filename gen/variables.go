package main

import (
  "log"
)

var C_MODULES = []string{
  "e_",
  "ecore_",
  "ecvt_",
  "edje_",
  "eet_",
  "efreet_",
  "eina_",
  "elm_",
  "ethumb_",
  "evas_",
  "eio_",
  "emotion_",
}

var MODULE_PRIORITY = map[string]int{
  "elm_": 90,
  "evas_": 30,
}

func modulePrefered(a, b string) bool {
  if MODULE_PRIORITY[a] == 0 { log.Fatal("set priority of ", a) }
  if MODULE_PRIORITY[b] == 0 { log.Fatal("set priority of ", b) }
  return MODULE_PRIORITY[a] > MODULE_PRIORITY[b]
}

var FUNCS_WITH_RETURN_PARAM = NewStringSet([]string{
  "elm_map_canvas_to_region_convert",
})

var NOT_SUPPORTED_TYPES = NewStringSet([]string{
  "va_list",
  "<function>",
  //TODO do it later
  "struct tm *",
  "const struct tm *",
})

var DISCARD_CONSTRUCT_FUNCS = NewStringSet([]string{
  "elm_win_util_standard_add",
})

var DISCARD_METHOD_FUNCS = NewStringSet([]string{
  // Evas_Callback_Type not supported
  "evas_object_event_callback_add",
  "evas_object_event_callback_priority_add",
  "evas_object_event_callback_del",
  "evas_object_event_callback_del_full",

  // Evas_Callback_Priority not supported
  "evas_object_smart_callback_priority_add",

  // varargs not supported
  "evas_object_box_option_property_set",

  //TODO I don't know their behaviours
  "elm_entry_filter_limit_size",
  "elm_entry_filter_accept_set",
  "elm_map_sources_get",

  //TODO manually implement these
  "elm_calendar_weekdays_names_get",

  // no use in go
  "elm_radio_value_pointer_set", // go value will not change
})

var PREFER_ENUM = map[string]string{
  "ASPECT_CONTROL_NONE": "EVAS_",
  "ASPECT_CONTROL_NEITHER": "EVAS_",
  "ASPECT_CONTROL_HORIZONTAL": "EVAS_",
  "ASPECT_CONTROL_VERTICAL": "EVAS_",
  "ASPECT_CONTROL_BOTH": "EVAS_",
  "OBJECT_TABLE_HOMOGENEOUS_NONE": "EVAS_",
  "OBJECT_TABLE_HOMOGENEOUS_TABLE": "EVAS_",
  "OBJECT_TABLE_HOMOGENEOUS_ITEM": "EVAS_",
  "LOAD_ERROR_NONE": "EVAS_",
  "LOAD_ERROR_GENERIC": "EVAS_",
  "LOAD_ERROR_DOES_NOT_EXIST": "EVAS_",
  "LOAD_ERROR_PERMISSION_DENIED": "EVAS_",
  "LOAD_ERROR_RESOURCE_ALLOCATION_FAILED": "EVAS_",
  "LOAD_ERROR_CORRUPT_FILE": "EVAS_",
  "LOAD_ERROR_UNKNOWN_FORMAT": "EVAS_",
  "INPUT_PANEL_LAYOUT_NORMAL": "ELM_",
  "INPUT_PANEL_LAYOUT_NUMBER": "ELM_",
  "INPUT_PANEL_LAYOUT_EMAIL": "ELM_",
  "INPUT_PANEL_LAYOUT_URL": "ELM_",
  "INPUT_PANEL_LAYOUT_PHONENUMBER": "ELM_",
  "INPUT_PANEL_LAYOUT_IP": "ELM_",
  "INPUT_PANEL_LAYOUT_MONTH": "ELM_",
  "INPUT_PANEL_LAYOUT_NUMBERONLY": "ELM_",
  "INPUT_PANEL_LAYOUT_INVALID": "ELM_",
  "INPUT_PANEL_LAYOUT_HEX": "ELM_",
  "INPUT_PANEL_LAYOUT_TERMINAL": "ELM_",
  "INPUT_PANEL_LAYOUT_PASSWORD": "ELM_",
  "INPUT_PANEL_LANG_AUTOMATIC": "ELM_",
  "INPUT_PANEL_LANG_ALPHABET": "ELM_",
  "INPUT_PANEL_RETURN_KEY_TYPE_DEFAULT": "ELM_",
  "INPUT_PANEL_RETURN_KEY_TYPE_DONE": "ELM_",
  "INPUT_PANEL_RETURN_KEY_TYPE_GO": "ELM_",
  "INPUT_PANEL_RETURN_KEY_TYPE_JOIN": "ELM_",
  "INPUT_PANEL_RETURN_KEY_TYPE_LOGIN": "ELM_",
  "INPUT_PANEL_RETURN_KEY_TYPE_NEXT": "ELM_",
  "INPUT_PANEL_RETURN_KEY_TYPE_SEARCH": "ELM_",
  "INPUT_PANEL_RETURN_KEY_TYPE_SEND": "ELM_",
}
