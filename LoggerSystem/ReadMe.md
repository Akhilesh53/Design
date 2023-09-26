Design a Logger 
======================

**A logger has following componenets.**

*Application -> Logger -> Log Types -> Sink*
                                   
LogTypes :INFO , DEBUG , ERROR

Sink :  FILE, QUEUE, DB


**We have to design in such a way**
-------------------------------------------
1) If some new log types has to be addded , the code should be extensible.
2) Same goes for SINK

-------------------------------------------

Now we got the objects

- Application
- Logger
- LogTypes
- Sink (Where the logs should be printed)

-------------------------------------------
Now relate these objects to some design pattern.
-------------------------------------------

1) Logger class must be defined only once. if defined multiple instance, there will be no sync in logs.
   Means logger object must follow singelton design pattern

2) Log Types: the functionality is whatsoever log level is coming that log type object should print the statement.
   Means once we will call the func, we don't care which handler is handling our request. They are following a chain.
   If log level is not info, check for error, if not error then check for debug.

3) Sink Types: the functionality is where so ever log type is defined to dump in, dump the logs there. 
   Means, sink types are observing if some level assigned to them comes, take that log and dump it.
   Folowing Observer pattern.


-----------------------------------
Try to define the high level func/ relationship for each objects.
-----------------------------------
- **Logger** :  
-----------------------------  
Association: **Logger -> HAS -> LogTypes**

- GetLoggerInstance()
- Info()
- Debug()
- Error()
    
-------------------------------
**LogTypes**:  
-------------------------------
- **ILogHandler Interface**
--------------------------------------------------
   - GetNext() ILogHandler
   - SetNext(next ILogHandler)
   - GetLevel() string
   - SetLevel(level LogType)
   - LogMessage(logType LogType, msg string)
   - AddSinkObserver(ISinkObserver)
	- RemoveSinkObserver(ISinkObserver)
	- NotifyInkObservers(mssg string)

Log handlers will also work as observables for sink observers. Whenever a log comes in, observable will notify
observers to dump the log
---------------------------------
- **Concrete Classes:**
---------------------------------
Association: **INFO/ DEBUG/ ERROR -> IS A -> LOG TYPE**
   - InfoHandler{}
   - DebugHandler{}  
   - ErrorHandler{}

---------------------------------
**Sink Types**
---------------------------------


