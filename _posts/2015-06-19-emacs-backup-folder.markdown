---
layout: bootstrap_post
title: How to Change Emacs Backup Folder
date: 2015-06-19 18:10:00
author: Oz Akan
abstract: One folder, out of sight, for all emacs backups
categories:
    - Tips
tags:
    - Systems
---

To configure emacs to store all backups under your home folder in `.emacs-backups` do this;

    # mkdir ~/.emacs-backups
    # emacs ~/.emacs

and the put at least the first four lines below in that file

    (setq backup-directory-alist 
            `((".*" . ,"~/.emacs-backups/")))
    (setq auto-save-file-name-transforms 
            `((".*" , "~/.emacs-backups/" t)))

    (setq delete-old-versions t
      kept-new-versions 6
      kept-old-versions 2
      version-control t)

    (setq-default indent-tabs-mode nil)
    (setq-default tab-width 4)
    (setq indent-line-function 'insert-tab)

Once saved, you will be free of backup mess that `vi` guys make fun of.

First section changes backup folder to `~/.emacs-backups`

Middle section changes number of backups to be 6.

Last section modifies tabs.