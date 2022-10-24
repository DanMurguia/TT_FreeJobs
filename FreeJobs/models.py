from django.db import models

# Create your models here.

class User(models.Model):
    userID = models.AutoField(primary_key = True)
    userName = models.CharField(max_length = 100)
    lastName = models.CharField(max_length = 100)
    #usename
    email = models.CharField(max_length = 100)
    password = models.CharField(max_length = 100)
    # code
    isActive = models.BoolField() #checar declaracion
    isAdmin = models.BoolField()
    isOffer = models.BoolField()
    createdAt = models.DateField()

    
class Profile(models.Model):
    birth = models.DateField()
    gender = models.CharField(max_length = 1)
    countryID = models.IntField() #FK
    image = models.CharField(max_length = 255)
    imageHeader = models.CharField(max_length = 255)
    title = models.CharField(max_length = 255)
    bio = models.CharField(max_length = 255)
    #likes
    #dislikes
    address = models.CharField(max_length = 255)
    phone = models.CharField(max_length = 100)
    email = models.CharField(max_length = 100)
    userID = models.AutoField(primary_key = True) #FK
    #levelID
    #sentimentalID

class Post(models.Model):
    postID = models.CharField(max_length = 500) #FK
    title = models.CharField(max_length = 500)
    content = models.CharField(max_length = 500)
    createdAt = models. DateField()

class Image(models.Model):
    imageID = models.AutoField(primary_key = True) #PK
    src = models.CharField(max_length = 500)
    title = models.CharField(max_length = 500)
    userID = models.CharField(max_length = 500) #FK
    albumID = models.CharField(max_length = 500) #FK
    createdAt = models.DateField()

class Heart(models.Model):
    heartID = models.AutoField(primary_key = True) #PK
    userID = title = models.CharField(max_length = 500) #FK
    date = title = models.DateField()

class Friend(models.Model): #Solo el oferente puede seguir
    friendID = models.AutoField(primary_key = True) #PK
    senderID = models.CharField(max_length = 500) #FK
    receptorID = models.CharField(max_length = 500) #FK

class Message(models.Models):
    messageID = models.AutoField(primary_key = True) #PK
    content = models.CharField(max_length = 500) #FK
    userID = models.CharField(max_length = 500) #FK
    conversationID = models.CharField(max_length = 500) #FK
    cretedAt = models.CharField(max_length = 500) #FK

class Conversarion(models.Model):
    converID = models.AutoField(primary_key = True) #PK
    senderID = models.CharField(max_length = 500) #FK
    receptorID = models.CharField(max_length = 500) #FK
    createdAt = models.DateField()

class Album(models.Model):
    albumID = models.AutoField(primary_key = True) #PK
    