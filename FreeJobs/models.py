from django.db import models

# Create your models here.

class User(models.Model):
    userID = models.AutoField(primary_key = True)
    userName = models.CharField(max_length = 100)
    lastName = models.CharField(max_length = 100)
    '''usename'''
    email = models.CharField(max_length = 100)
    password = models.CharField(max_length = 100)
    '''code'''
    isActive = models.BoolField() '''checar declaracion'''
    isAdmin = models.BoolField()
    isOffer = models.BoolField()
    createdAt = models.DateField()

class Profile(models.Model):
    birth = models.DateField()
    gender = models.CharField(max_length = 1)
    