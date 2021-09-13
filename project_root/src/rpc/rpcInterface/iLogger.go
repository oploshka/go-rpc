package rpcInterface


// Logger ...
//
// Сообщение ДОЛЖНО быть строкой или объектом, реализующим __toString().
//
// Сообщение МОЖЕТ содержать заполнители в форме: {foo} где foo
// будут заменены контекстными данными в ключе "foo".
//
// Контекстный массив может содержать произвольные данные, единственное предположение, что
// может быть сделано разработчиками, если указан экземпляр исключения
// для создания трассировки стека она ДОЛЖНА быть в ключе с именем "исключение".
//
type Logger interface {
    
    // Info
    // Interesting events.
    // Example: User logs in, SQL logs.
    Info (args ...interface{})
    
    // Warning
    // Exceptional occurrences that are not errors.
    // Example: Use of deprecated APIs, poor use of an API, undesirable things that are not necessarily wrong.
    Warning (args ...interface{})
    
    // Error
    // Runtime errors that do not require immediate action but should typically be logged and monitored.
    Error (args ...interface{})
    
    // Debug
    // Detailed debug information.
    Debug (args ...interface{})
    
    /*
    
    // Emergency System is unusable. Система непригодна для использования.
    Emergency()
    
    // Alert
    // Action must be taken immediately.
    // Example: Entire website down, database unavailable, etc. This should trigger the SMS alerts and wake you up.
    Alert()

    // Critical
    // Critical conditions.
    // Example: Application component unavailable, unexpected exception.
    Critical ()
    
    
    
    // Notice
    // Normal but significant events.
    Notice ()
    
    

    // Log - Logs with an arbitrary level.
    // Log($level, $message, array $context = array());
    */
}


