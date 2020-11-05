'기본 파라미터 생성
Function DEF_PARAM(jsonItems)
    Dim excelRange As Range
    'Dim jsonItems As New Collection
    Dim jsonDictionary As New Dictionary

    Dim r, c
    r = 7
    c = 12

    ' 셀을 선택
    ' Range("L7").Select
    ' 빈 셀에 도달하면 루프 수행이 중지되도록 설정합니다.
    Do Until IsEmpty(Cells(r, c).value)
        ' 여기에 코드를 삽입합니다.
        jsonDictionary("key") = Cells(r, c)
        jsonDictionary("value") = Cells(r + 2, c)
        jsonDictionary("description") = Cells(r + 1, c)
        jsonItems.Add jsonDictionary
        Set jsonDictionary = Nothing
        ' 현재 위치에서 1행 아래로 내려갑니다.
        ' ActiveCell.Offset(0, 1).Select
        c = c + 1
    Loop
    Set DEF_PARAM = jsonItems
    'Debug.Print JsonConverter.ConvertToJson(jsonItems, Whitespace:=3)
End Function

'Postman FORMDATA 생성
Function MAKE_FORMDATA(r, c)
    Dim jsonFormDataDictionary As New Dictionary
    Dim jsonItems As New Collection

    ' 빈 셀에 도달하면 루프 수행이 중지되도록 설정합니다.
    Do Until IsEmpty(Cells(r, c).value)
        ' 여기에 코드를 삽입합니다.
        jsonFormDataDictionary("key") = Cells(r, 4)
        jsonFormDataDictionary("value") = Cells(r, 6)
        jsonFormDataDictionary("description") = Cells(r, 5)
        jsonFormDataDictionary("disabled") = Cells(r, 7)
        jsonItems.Add jsonFormDataDictionary
        Set jsonFormDataDictionary = Nothing
        ' 현재 위치에서 1행 아래로 내려갑니다.
        r = r + 1
    Loop

    ' 기본 파라미터 생성
    DEF_PARAM jsonItems

    'jsonFormDataDictionary("formdata") = jsonItems

    Set MAKE_FORMDATA = jsonItems
    'Debug.Print JsonConverter.ConvertToJson(jsonFormDataDictionary, Whitespace:=3)
End Function

'Postman URL 생성
Function MAKE_URL(r, c)
    Dim jsonUrlDictionary As New Dictionary

    Dim jsonUrl As New Collection
    Dim strUrl() As String
    strUrl = Split(Range("L3").value, ".")

    For i = LBound(strUrl) To UBound(strUrl)
        If Len(strUrl(i)) > 0 Then
            jsonUrl.Add strUrl(i)
        End If
    Next i


    Dim jsonPath As New Collection
    Dim strPath() As String
    strPath = Split(Cells(r, c).value, "/")

    For i = LBound(strPath) To UBound(strPath)
        If Len(strPath(i)) > 0 Then
            jsonPath.Add strPath(i)
        End If
    Next i

    jsonUrlDictionary("raw") = "https://" & Range("L3") & Cells(r, c).value
    jsonUrlDictionary("protocol") = "https"
    jsonUrlDictionary("host") = jsonUrl
    jsonUrlDictionary("path") = jsonPath

    'jsonFormDataDictionary("url") = jsonDictionary

    Set MAKE_URL = jsonUrlDictionary
    'Debug.Print JsonConverter.ConvertToJson(jsonFormDataDictionary, Whitespace:=3)
End Function

'Postman API 생성
Function MAKE_API(r, c)
    Dim startR, startC
    Dim jsonItemDictionary As New Dictionary

    startR = r
    startC = c

    ' 셀을 선택
    ' Cells(r, c).Select
    jsonItemDictionary("name") = Cells(r, 3)

    Dim jsonBody As New Dictionary
    jsonBody("mode") = "formdata"
    jsonBody("formdata") = MAKE_FORMDATA(r, c + 2)
    
    Dim jsonRequest As New Dictionary
    jsonRequest("method") = "POST"
    jsonRequest("body") = jsonBody
    jsonRequest("url") = MAKE_URL(startR, startC)

    jsonItemDictionary("request") = jsonRequest

    Set MAKE_API = jsonItemDictionary
    'Debug.Print JsonConverter.ConvertToJson(jsonItemDictionary, Whitespace:=3)
End Function

'응답용 데이터 생성
Function MAKE_RESPONCE(r, c)
    Dim startR, startC
    Dim apiName

    startR = r
    startC = c

    apiName = Cells(startR, startC)
    apiName = Replace(apiName, "/", "_")
    apiName = UCase(apiName)
    apiName = Right(apiName, Len(apiName) - 1)

    MAKE_RESPONCE = apiName
    'Worksheets("response").Cells(startR, startC).Value = apiName
End Function

Sub MAKE_IOS_HEADER(className)
    Dim strTemplete
    Dim templete() As String
    Dim strHeader
    Dim fileName
    
    strTemplete = Worksheets("ios_template").Cells(3, 3).value
    
    templete = Split(strTemplete, "|")
    strHeader = templete(0) & className & templete(1)
    fileName = "ios/" & className & ".h"
    
    SaveFile strHeader, fileName
    Debug.Print fileName
End Sub

Sub MAKE_IOS_CLASS(className)
    MakeFolderIfNotExist (ThisWorkbook.Path & "/ios")
    MAKE_IOS_HEADER (className)
End Sub

'응답값 분석
Sub PARSE_RESPONSE()
    Dim startR, startC, OffsetR
    Dim strJson
    Dim jsonData As Dictionary

    startR = 3
    startC = 3

    Do Until IsEmpty(Worksheets("response").Cells(startR, startC).value)

        strJson = ReadFile(Worksheets("response").Cells(startR, startC).value)
        If Len(strJson) > 0 Then
            MAKE_IOS_CLASS ("DATA_" & Worksheets("response").Cells(startR, startC).value)
            'Debug.Print strJson
            Set jsonData = JsonConverter.ParseJson(strJson)
            
            Dim key As Variant
            Dim value As Variant
            For Each key In jsonData.Keys
              Debug.Print key & ", " & TypeName(jsonData(key))
              
              If TypeName(jsonData(key)) = "Collection" Then
              ElseIf TypeName(jsonData(key)) = "Dictionary" Then
              Else
                value = jsonData(key)
                Debug.Print key & ", " & value
              End If
            Next key
        End If
        
        startR = startR + 1
    Loop
End Sub


Sub start()
    'TrimCell

    Dim jsonInfoDictionary As New Dictionary
    Dim jsonItems As New Collection
    Dim startR, startC, OffsetR
    '응답 테이블 관련 변수
    Dim responseR

    startR = 3
    startC = 2
    OffsetR = 1
    responseR = 3

    '응답 테이블 삭제
    Range("response!B3:C2000").Clear

    Dim jsonInfo As New Dictionary
    jsonInfo("name") = "api-moolban"
    jsonInfo("schema") = "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"


    Do Until IsEmpty(Cells(startR, startC).value)

        OffsetR = 1
        If Len(Cells(startR, 5).value) = 0 And Len(Cells(startR + 1, 5).value) = 0 Then
            OffsetR = OffsetR + 1
        End If

        '응답 테이블 처리
        Worksheets("response").Cells(responseR, 2).value = Cells(startR, startC).value
        Worksheets("response").Cells(responseR, 3).value = MAKE_RESPONCE(startR, startC)
        responseR = responseR + 1
        '응답 테이블 처리 끝

        'Debug.Print "api [" & startR & "," & startC & "]"
        Dim jsonFormDataDictionary
        Set jsonFormDataDictionary = MAKE_API(startR, startC)
        
        jsonItems.Add jsonFormDataDictionary
        startR = startR + OffsetR
    Loop

    jsonInfoDictionary("info") = jsonInfo
    jsonInfoDictionary("item") = jsonItems

    'ThisWorkbook.Sheets("result").Cells(1, 1) = JsonConverter.ConvertToJson(jsonInfoDictionary, Whitespace:=3)
    'Debug.Print JsonConverter.ConvertToJson(jsonItems, Whitespace:=3)
    SaveFile JsonConverter.ConvertToJson(jsonInfoDictionary, Whitespace:=3)


    PARSE_RESPONSE
End Sub

Sub SaveFile(strValue, Optional fileName = "")
    Dim outputFilePath, outputFile
    Dim strfullpath As String
    
    If Len(fileName) > 0 Then
        strfullpath = ThisWorkbook.Path & "/" & fileName
    Else
        strfullpath = ThisWorkbook.Path & "/api-moolban.postman_collection.json"
    End If

    Debug.Print strfullpath
    
    outputFile = FreeFile
    Open strfullpath For Output As #outputFile
        Print #outputFile, strValue
        'Print #outputFile, ThisWorkbook.Sheets("result").Cells(1, 1).Value
    Close #outputFile

End Sub

Function MakeFolderIfNotExist(Folderstring As String)
'Ron de Bruin, 22-June-2015
' http://www.rondebruin.nl/mac/mac010.htm
    Dim ScriptToMakeFolder As String
    Dim str As String
    If Val(Application.Version) < 15 Then
        ScriptToMakeFolder = "tell application " & Chr(34) & _
                             "Finder" & Chr(34) & Chr(13)
        ScriptToMakeFolder = ScriptToMakeFolder & _
                "do shell script ""mkdir -p "" & quoted form of posix path of (" & _
                        Chr(34) & Folderstring & Chr(34) & ")" & Chr(13)
        ScriptToMakeFolder = ScriptToMakeFolder & "end tell"
        On Error Resume Next
        MacScript (ScriptToMakeFolder)
        On Error GoTo 0

    Else
        str = MacScript("return POSIX path of (" & _
                        Chr(34) & Folderstring & Chr(34) & ")")
        MkDir str
    End If
End Function



Function ReadFile(strValue)
    Dim FileNum As Integer
    Dim DataLine As String, MyPath As String

    FileNum = FreeFile()
    MyPath = ThisWorkbook.Path & "/" & strValue & ".html"

On Error GoTo Catch
    Open MyPath For Input As #FileNum
    While Not EOF(FileNum)
        Line Input #FileNum, DataLine ' read in data 1 line at a time
        'Debug.Print DataLine
    Wend
    ReadFile = DataLine

    GoTo Finally
Catch:
    ReadFile = ""
Finally:
End Function


Sub TrimCell()
    Dim cell As Range
    For Each cell In ActiveSheet.UsedRange.SpecialCells(xlCellTypeConstants)
    cell = WorksheetFunction.Trim(cell)
    Next cell
End Sub


